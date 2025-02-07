// Copyright 2021 Comcast Cable Communications Management, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docs

// swagger:route POST /v1/orgs/{orgid}/applications/{appid}/routes routes postRoute
// Adds a new route to the routing table or updates an existing route. Route ID can be given in the body. If it is omitted a hash will be calculates and used instead.
// responses:
//   200: RouteResponse
//   500: RouteErrorResponse
