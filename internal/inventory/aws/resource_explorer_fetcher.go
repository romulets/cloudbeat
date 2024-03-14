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
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/elastic/elastic-agent-libs/logp"

	"github.com/elastic/cloudbeat/internal/inventory"
	"github.com/elastic/cloudbeat/internal/resources/providers/awslib"
	"github.com/elastic/cloudbeat/internal/resources/providers/awslib/resourceexplorer"
)

// var allowedAssetTypes []string

type resourcesExplorerProvider interface {
	SearchAll(ctx context.Context) ([]resourceexplorer.Resource, error)
}

type ResourceExplorerFetcher struct {
	logger   *logp.Logger
	provider resourcesExplorerProvider
}

func NewResourceExplorerFetcher(logger *logp.Logger, cfg aws.Config) inventory.AssetFetcher {
	provider := resourceexplorer.NewProvider(logger, cfg, &awslib.MultiRegionClientFactory[resourceexplorer.Client]{})
	return &ResourceExplorerFetcher{
		logger:   logger,
		provider: provider,
	}
}

func (a *ResourceExplorerFetcher) Fetch(ctx context.Context, assetChannel chan<- inventory.AssetEvent) error {
	resources, err := a.provider.SearchAll(ctx)
	if err != nil {
		return fmt.Errorf("could not list ec2 instances (%w)", err)
	}

	var errAgg error
	for _, resource := range resources {
		if resource.Arn == nil {
			continue
		}

		parsedArn, inErr := arn.Parse(*resource.Arn)
		if inErr != nil {
			errAgg = errors.Join(errAgg, inErr)
			continue
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
			*resource.Arn,
			"",
			inventory.WithRawAsset(resource),
			inventory.WithMetadata(inventory.AssetMetadata{
				SourceAPI: "resource_explorer",
			}),
			inventory.WithCloud(inventory.AssetCloud{
				Provider:  inventory.AwsCloudProvider,
				Region:    parsedArn.Region,
				AccountId: parsedArn.AccountID,
			}),
		)
	}

	return errAgg
}
