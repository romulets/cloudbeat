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

package resourcetags

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	resourcegroup "github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/elastic/cloudbeat/internal/resources/providers/awslib"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/samber/lo"
)

type Resource types.ResourceTagMapping

type Client interface {
	GetResources(ctx context.Context, params *resourcegroup.GetResourcesInput, optFns ...func(*resourcegroup.Options)) (*resourcegroup.GetResourcesOutput, error)
}

type Provider struct {
	log     *logp.Logger
	clients map[string]Client
}

func NewProvider(log *logp.Logger, cfg aws.Config, factory awslib.CrossRegionFactory[Client]) *Provider {
	f := func(cfg aws.Config) Client {
		return resourcegroup.NewFromConfig(cfg)
	}
	m := factory.NewMultiRegionClients(awslib.AllRegionSelector(), cfg, f, log)

	return &Provider{
		log:     log,
		clients: m.GetMultiRegionsClientMap(),
	}
}

func (p *Provider) GetResources(ctx context.Context) ([]Resource, error) {
	resources, err := awslib.MultiRegionFetch(ctx, p.clients, func(ctx context.Context, _ string, c Client) ([]Resource, error) {
		params := resourcegroup.GetResourcesInput{}
		allResources := make([]Resource, 0)

		for {
			output, err := c.GetResources(ctx, &params)
			if err != nil {
				return nil, err
			}

			if output == nil || len(output.ResourceTagMappingList) == 0 {
				break
			}

			for _, r := range output.ResourceTagMappingList {
				allResources = append(allResources, Resource(r))
			}

			if output.PaginationToken == nil || *output.PaginationToken == "" {
				break
			}
			params.PaginationToken = output.PaginationToken
		}

		return allResources, nil
	})

	if err != nil {
		return nil, err
	}

	return lo.Flatten(resources), nil
}
