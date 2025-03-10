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

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	ecrClient "github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"

	"github.com/elastic/cloudbeat/internal/infra/clog"
	"github.com/elastic/cloudbeat/internal/resources/providers/awslib"
)

type Repository types.Repository
type Provider struct {
	clients map[string]Client
}

type RepositoryDescriber interface {
	DescribeRepositories(ctx context.Context, repoNames []string, region string) ([]types.Repository, error)
}

type Client interface {
	DescribeRepositories(ctx context.Context, params *ecrClient.DescribeRepositoriesInput, optFns ...func(*ecrClient.Options)) (*ecrClient.DescribeRepositoriesOutput, error)
}

func NewEcrProvider(ctx context.Context, log *clog.Logger, cfg aws.Config, factory awslib.CrossRegionFactory[Client]) *Provider {
	f := func(cfg aws.Config) Client {
		return ecrClient.NewFromConfig(cfg)
	}
	m := factory.NewMultiRegionClients(ctx, awslib.AllRegionSelector(), cfg, f, log)

	return &Provider{clients: m.GetMultiRegionsClientMap()}
}
