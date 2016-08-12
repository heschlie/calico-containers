// Copyright (c) 2016 Tigera, Inc. All rights reserved.

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

package commands

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	"github.com/tigera/libcalico-go/calicoctl/resourcemgr"
	"github.com/tigera/libcalico-go/lib/api"
	"github.com/tigera/libcalico-go/lib/api/unversioned"
	"github.com/tigera/libcalico-go/lib/client"
	"github.com/tigera/libcalico-go/lib/net"
)

// Create a new CalicoClient using connection information in the specified
// filename (if it exists), dropping back to environment variables for any
// parameter not loaded from file.
func newClient(cf string) (*client.Client, error) {
	if _, err := os.Stat(cf); err != nil {
		glog.V(2).Infof("Config file cannot be read - reading config from environment")
		cf = ""
	}

	cfg, err := client.LoadClientConfig(cf)
	if err != nil {
		return nil, err
	}
	glog.V(2).Infof("Loaded client config: type=%v %#v", cfg.BackendType, cfg.BackendConfig)

	c, err := client.New(*cfg)
	if err != nil {
		return nil, err
	}

	return c, err
}
