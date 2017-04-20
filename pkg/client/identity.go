// Copyright 2016-2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"github.com/polyverse-security/cilium/api/v1/client/policy"
	"github.com/polyverse-security/cilium/api/v1/models"
)

// IdentityGet returns a security identity.
func (c *Client) IdentityGet(id string) (*models.Identity, error) {
	params := policy.NewGetIdentityIDParams().WithID(id)

	resp, err := c.Policy.GetIdentityID(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
