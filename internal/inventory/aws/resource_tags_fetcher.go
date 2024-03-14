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

package aws

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/elastic/elastic-agent-libs/logp"

	"github.com/elastic/cloudbeat/internal/inventory"
	"github.com/elastic/cloudbeat/internal/resources/providers/awslib"
	"github.com/elastic/cloudbeat/internal/resources/providers/awslib/resourcetags"
	"github.com/elastic/cloudbeat/internal/resources/utils/pointers"
)

type resourceTagsProvider interface {
	GetResources(ctx context.Context) ([]resourcetags.Resource, error)
}

type ResourceTagsFetcher struct {
	logger   *logp.Logger
	provider resourceTagsProvider
}

func NewResourceTagsFetcher(logger *logp.Logger, cfg aws.Config) inventory.AssetFetcher {
	provider := resourcetags.NewProvider(logger, cfg, &awslib.MultiRegionClientFactory[resourcetags.Client]{})
	return &ResourceTagsFetcher{
		logger:   logger,
		provider: provider,
	}
}

func (a *ResourceTagsFetcher) Fetch(ctx context.Context, assetChannel chan<- inventory.AssetEvent) error {
	resources, err := a.provider.GetResources(ctx)
	if err != nil {
		return fmt.Errorf("could not list ec2 instances (%w)", err)
	}

	var errAgg error
	for _, resource := range resources {
		if resource.ResourceARN == nil {
			continue
		}

		parsedArn, inErr := arn.Parse(*resource.ResourceARN)
		if inErr != nil {
			errAgg = errors.Join(errAgg, inErr)
			continue
		}

		name := getNameFromTags(resource.Tags)
		tags := make(map[string]string, len(resource.Tags))
		for _, t := range resource.Tags {
			if t.Key == nil {
				continue
			}

			tags[*t.Key] = pointers.Deref(t.Value)
		}

		subType := parsedArn.Service
		resourceType := ""

		if strings.Contains(parsedArn.Resource, "/") {
			resourceType = strings.Split(parsedArn.Resource, "/")[0]
		}

		if strings.Contains(parsedArn.Resource, ":") {
			resourceType = strings.Split(resourceType, ":")[0]
		}

		if resourceType != "" {
			subType += "-" + resourceType
		}

		assetChannel <- inventory.NewAssetEvent(
			inventory.AssetClassification{
				Category:    inventory.CategoryInfrastructure,
				SubCategory: inventory.SubCategoryCompute,
				Type:        inventory.TypeVirtualMachine,
				SubType:     subType,
			},
			*resource.ResourceARN,
			name,
			inventory.WithMetadata(inventory.AssetMetadata{
				SourceAPI: "resource_tags",
			}),
			inventory.WithTags(tags),
			inventory.WithCloud(inventory.AssetCloud{
				Provider:  inventory.AwsCloudProvider,
				Region:    parsedArn.Region,
				AccountId: parsedArn.AccountID,
			}),
		)
	}

	return errAgg
}

func getNameFromTags(tags []types.Tag) string {
	for _, tag := range tags {
		if strings.ToLower(pointers.Deref(tag.Key)) == "name" {
			return pointers.Deref(tag.Value)
		}
	}
	return ""
}
