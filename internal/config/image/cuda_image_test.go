/**
# Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
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

package image

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMajorMinorVersionValid(t *testing.T) {
	var tests = []struct {
		version  string
		expected string
	}{
		{"0", "0.0"},
		{"8", "8.0"},
		{"7.5", "7.5"},
		{"9.0.116", "9.0"},
		{"4294967295.4294967295.4294967295", "4294967295.4294967295"},
		{"v11.6", "11.6"},
	}
	for _, c := range tests {
		t.Run(c.version, func(t *testing.T) {
			version, err := parseMajorMinorVersion(c.version)

			require.NoError(t, err)
			require.Equal(t, c.expected, version)
		})
	}
}

func TestParseMajorMinorVersionInvalid(t *testing.T) {
	var tests = []string{
		"foo",
		"foo.5.10",
		"9.0.116.50",
		"9.0.116foo",
		"7.foo",
		"9.0.bar",
		"9.4294967296",
		"9.0.116.",
		"9..0",
		"9.",
		".5.10",
		"-9",
		"+9",
		"-9.1.116",
		"-9.-1.-116",
	}
	for _, c := range tests {
		t.Run(c, func(t *testing.T) {
			_, err := parseMajorMinorVersion(c)
			require.Error(t, err)
		})
	}
}