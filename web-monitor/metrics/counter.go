// Copyright (c) 2018 Baidu, Inc.
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

package metrics

import (
	"sync/atomic"
)

// Counter is a cumulative metric that represents a single monotonically increasing counter
// whose value can only increase or be reset to zero on restart
type Counter uint64

// Inc increases counter
func (c *Counter) Inc(delta uint) {
	if c == nil {
		return
	}
	atomic.AddUint64((*uint64)(c), uint64(delta))
}

// Get gets counter
func (c *Counter) Get() int64 {
	if c == nil {
		return 0
	}
	return int64(atomic.LoadUint64((*uint64)(c)))
}

func (c *Counter) Type() string {
	return TypeCounter
}
