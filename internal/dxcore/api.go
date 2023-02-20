/**
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package dxcore

import (
	"github.com/NVIDIA/go-nvml/pkg/dl"
)

const (
	libraryName      = "libdxcore.so"
	libraryLoadFlags = dl.RTLD_LAZY | dl.RTLD_GLOBAL
)

// dxcore stores a reference the dxcore dynamic library
var dxcore *context

// Init initializes the dxcore dynamic library
func Init() error {
	c, err := initContext()
	if err != nil {
		return err
	}
	dxcore = c
	return nil
}

// Shutdown closes the dxcore dynamic library
func Shutdown() error {
	if dxcore != nil && dxcore.initialized != 0 {
		dxcore.deinitContext()
	}
	return nil
}

// GetDriverStorePaths returns the list of driver store paths
func GetDriverStorePaths() []string {
	var paths []string
	for i := 0; i < dxcore.getAdapterCount(); i++ {
		adapter := dxcore.getAdapter(i)
		paths = append(paths, adapter.getDriverStorePath())
	}

	return paths
}
