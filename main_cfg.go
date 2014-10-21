//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the
//  License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an "AS
//  IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
//  express or implied. See the License for the specific language
//  governing permissions and limitations under the License.

package main

import (
	"fmt"
	"os"
)

func MainCfg(provider, dataDir string) (Cfg, error) {
	// TODO: One day, the default cfg provider should not be simple
	if provider == "" || provider == "simple" {
		return MainCfgSimple(provider, dataDir)
	}
	return nil, fmt.Errorf("error: unsupported cfg provider: %s", provider)
}

func MainCfgSimple(provider, dataDir string) (Cfg, error) {
	cfgPath := dataDir + string(os.PathSeparator) + "cbft.cfg"
	cfgPathExists := false
	if _, err := os.Stat(cfgPath); err == nil {
		cfgPathExists = true
	}

	cfg := NewCfgSimple(cfgPath)
	if cfgPathExists {
		err := cfg.Load()
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
