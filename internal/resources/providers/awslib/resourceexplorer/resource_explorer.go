// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package resourceexplorer

import (
	"context"
	"slices"

	"github.com/aws/aws-sdk-go-v2/aws"
	resourceexplorer "github.com/aws/aws-sdk-go-v2/service/resourceexplorer2"
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/samber/lo"

	"github.com/elastic/cloudbeat/internal/resources/providers/awslib"
	"github.com/elastic/cloudbeat/internal/resources/utils/pointers"
)

type Resource types.Resource

type Client interface {
	Search(ctx context.Context, params *resourceexplorer.SearchInput, optFns ...func(*resourceexplorer.Options)) (*resourceexplorer.SearchOutput, error)
}

type RegionSelector struct{}

type Provider struct {
	log     *logp.Logger
	clients map[string]Client
}

func NewProvider(log *logp.Logger, cfg aws.Config, factory awslib.CrossRegionFactory[Client]) *Provider {
	f := func(cfg aws.Config) Client {
		return resourceexplorer.NewFromConfig(cfg)
	}
	m := factory.NewMultiRegionClients(&RegionSelector{}, cfg, f, log)

	return &Provider{
		log:     log,
		clients: m.GetMultiRegionsClientMap(),
	}
}

func (r *RegionSelector) Regions(ctx context.Context, cfg aws.Config) ([]string, error) {
	c := resourceexplorer.NewFromConfig(cfg)

	params := resourceexplorer.ListIndexesInput{
		Type: types.IndexTypeAggregator,
	}

	regions := make([]string, 0, 100)
	for {
		output, err := c.ListIndexes(ctx, &params)
		if err != nil {
			return nil, err
		}

		if output == nil || len(output.Indexes) == 0 {
			break
		}

		for _, i := range output.Indexes {
			if i.Region == nil || *i.Region == "" || slices.Contains(regions, *i.Region) {
				continue
			}

			regions = append(regions, *i.Region)
		}

		if output.NextToken == nil || *output.NextToken == "" {
			break
		}
		params.NextToken = output.NextToken
	}

	return regions, nil
}

func (p *Provider) SearchAll(ctx context.Context) ([]Resource, error) {
	resources, err := awslib.MultiRegionFetch(ctx, p.clients, func(ctx context.Context, region string, c Client) ([]Resource, error) {
		p.log.Infof("Region '%s'", region)
		params := resourceexplorer.SearchInput{
			QueryString: pointers.Ref(""),
		}

		allResources := make([]Resource, 0)

		for {
			output, err := c.Search(ctx, &params)
			if err != nil {
				return nil, err
			}

			if output == nil {
				break
			}

			for _, r := range output.Resources {
				allResources = append(allResources, Resource(r))
			}

			if output.NextToken == nil || *output.NextToken == "" {
				break
			}
			params.NextToken = output.NextToken
		}

		return allResources, nil
	})

	if err != nil {
		return nil, err
	}

	return lo.Flatten(resources), nil
}
