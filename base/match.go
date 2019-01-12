/*
 *
 *
 *  * Copyright 2019 koalaone@163.com
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *       http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package base

import (
	"reflect"

	"github.com/koalaone/mock/interfaces"
)

type Match struct {
	Report
	values  map[string]int
	matches []string
}

func NewMatch(matches []string, names []string) *Match {
	return makeMatch(matches, names)
}

func makeMatch(matches []string, names []string) *Match {
	if matches == nil {
		matches = []string{}
	}
	values := map[string]int{}
	for id, name := range names {
		if name != "" {
			values[name] = id
		}
	}
	return &Match{matches: matches, values: values}
}

func (mh *Match) Raw() []string {
	return mh.matches
}

func (mh *Match) Length() interfaces.IInteger {
	return NewInteger(len(mh.matches))
}

func (mh *Match) Index(index int) interfaces.IString {
	if index < 0 || index >= len(mh.matches) {
		mh.LogError("Match Index error: out of range")
		return NewString("")
	}

	return NewString(mh.matches[index])
}

func (mh *Match) Name(value string) interfaces.IString {
	index, ok := mh.values[value]
	if !ok {
		mh.LogError("Match Name not found: %v", value)
		return NewString("")
	}

	return mh.Index(index)
}

func (mh *Match) EqualEmpty() interfaces.IMatch {
	if len(mh.matches) != 0 {
		mh.LogError("Match EqualEmpty empty")
	}

	return mh
}

func (mh *Match) NotEqualEmpty() interfaces.IMatch {
	if len(mh.matches) == 0 {
		mh.LogError("Match EqualEmpty empty")
	}

	return mh
}

func (mh *Match) EqualValues(values ...string) interfaces.IMatch {
	if values == nil {
		values = []string{}
	}

	if !reflect.DeepEqual(values, mh.values) {
		mh.LogError("Match EqualValues empty")
	}

	return mh
}

func (mh *Match) NotEqualValues(values ...string) interfaces.IMatch {
	if values == nil {
		values = []string{}
	}

	if reflect.DeepEqual(values, mh.values) {
		mh.LogError("Match NotEqualValues empty")
	}

	return mh
}
