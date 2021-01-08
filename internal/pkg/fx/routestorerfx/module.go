// Copyright 2020 Comcast Cable Communications Management, LLC
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

package routestorerfx

import (
	"errors"

	"github.com/xmidt-org/ears/internal/pkg/app"
	"github.com/xmidt-org/ears/internal/pkg/db"
	"github.com/xmidt-org/ears/pkg/route"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		ProvideRouteStorer,
	),
)

type StorageIn struct {
	fx.In

	Config app.Config
}

type StorageOut struct {
	fx.Out

	RouteStorer route.RouteStorer `name:"RouteStorer"`
}

func ProvideRouteStorer(in StorageIn) (StorageOut, error) {
	out := StorageOut{}

	stroageType := in.Config.GetString("ears.storageType")

	if stroageType == "inmemory" {
		out.RouteStorer = db.NewInMemoryRouteStorer(in.Config)
	} else {
		return out, errors.New("usupported storage type " + stroageType)
	}

	return out, nil
}
