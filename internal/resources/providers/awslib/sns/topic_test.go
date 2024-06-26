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

package sns

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloudbeat/internal/resources/utils/pointers"
)

func TestSNS_TopicInfo(t *testing.T) {
	tests := []struct {
		name         string
		resource     *TopicInfo
		expectedName string
		expectedArn  string
	}{
		{
			name: "everything's perfect and it just works",
			resource: &TopicInfo{
				Topic: types.Topic{
					TopicArn: pointers.Ref("aws:arn:hi"),
				},
			},
			expectedName: "hi",
			expectedArn:  "aws:arn:hi",
		},
		{
			name: "no ARN means no name",
			resource: &TopicInfo{
				Topic: types.Topic{
					TopicArn: pointers.Ref(""),
				},
			},
			expectedName: "",
			expectedArn:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName := tt.resource.GetResourceName()
			gotArn := tt.resource.GetResourceArn()

			assert.Equal(t, tt.expectedName, gotName)
			assert.Equal(t, tt.expectedArn, gotArn)
		})
	}
}
