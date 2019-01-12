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
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/koalaone/mock/interfaces"
)

type String struct {
	Report
	value string
}

func NewString(value string) *String {
	return &String{value: value}
}

func (sg *String) Raw() string {
	return sg.value
}

func (sg *String) Path(path string) interfaces.IValue {
	newValue, err := convertPath(sg.value, path)
	if err != nil {
		sg.LogError("String Path error:%v", err)
		return &Value{value: sg.value}
	}

	return newValue
}

func (sg *String) Schema(value interface{}) interfaces.IString {
	err := checkSchema(sg.value, value)
	if err != nil {
		sg.LogError("Array Schema error:%v", err)
	}

	return sg
}

func (sg *String) Equal(value interface{}) interfaces.IString {
	if !reflect.DeepEqual(sg.value, value) {
		sg.LogError("String Equal convert error:%v, target:%v", sg.value, value)
	}

	return sg
}

func (sg *String) NotEqual(value interface{}) interfaces.IString {
	if reflect.DeepEqual(sg.value, value) {
		sg.LogError("String NotEqual convert error:%v, target:%v", sg.value, value)
	}

	return sg
}

func (sg *String) EqualFold(value string) interfaces.IString {
	if !strings.EqualFold(sg.value, value) {
		sg.LogError("String EqualFold convert error:%v, target:%v", sg.value, value)
	}

	return sg
}

func (sg *String) NotEqualFold(value string) interfaces.IString {
	if strings.EqualFold(sg.value, value) {
		sg.LogError("String NotEqualFold convert error:%v, target:%v", sg.value, value)
	}

	return sg
}

func (sg *String) Length() interfaces.IInteger {
	return NewInteger(len(sg.value))
}

func (sg *String) Datetime(layout string) interfaces.IDatetime {
	var locTime time.Time
	var err error
	if layout != "" {
		locTime, err = time.Parse(layout, sg.value)
		if err != nil {
			sg.LogError("String Datetime convert error:%v", sg.value)
		}
	} else {
		locTime, err = http.ParseTime(sg.value)
		if err != nil {
			sg.LogError("String Datetime convert error:%v", sg.value)
		}
	}

	return NewDatetime(locTime)
}

func (sg *String) EqualEmpty() interfaces.IString {
	return sg.Equal("")
}

func (sg *String) NotEqualEmpty() interfaces.IString {
	return sg.NotEqual("")
}

func (sg *String) Contains(value string) interfaces.IString {
	if !strings.Contains(sg.value, value) {
		sg.LogError("String Contains convert error:%v, target:%v", sg.value, value)
	}

	return sg
}

func (sg *String) NotContains(value string) interfaces.IString {
	if strings.Contains(sg.value, value) {
		sg.LogError("String NotContains convert error:%v, target:%v", sg.value, value)
	}

	return sg
}

func (sg *String) ContainFold(value string) interfaces.IString {
	if !strings.Contains(strings.ToLower(sg.value), strings.ToLower(value)) {
		sg.LogError("String ContainFold convert error:%v, target:%v", sg.value, value)
	}

	return sg
}

func (sg *String) NotContainFold(value string) interfaces.IString {
	if strings.Contains(strings.ToLower(sg.value), strings.ToLower(value)) {
		sg.LogError("String NotContainFold convert error:%v, target:%v", sg.value, value)
	}

	return sg
}

func (sg *String) Match(expr string) interfaces.IMatch {
	ep, err := regexp.Compile(expr)
	if err != nil {
		sg.LogError("String Match convert error:%v , expr:%v", sg.value, expr)

	}

	matches := ep.FindStringSubmatch(sg.value)
	if matches == nil {
		return makeMatch(nil, nil)
	}

	return makeMatch(matches, ep.SubexpNames())
}

func (sg *String) MatchMulti(expr string) []interfaces.IMatch {
	result := make([]interfaces.IMatch, 0)
	ep, err := regexp.Compile(expr)
	if err != nil {
		sg.LogError("String MatchMulti convert error:%v , expr:%v", sg.value, expr)
		return result
	}

	matches := ep.FindAllStringSubmatch(sg.value, -1)
	if matches == nil {
		sg.LogError("String MatchMulti convert error:%v , expr:%v", sg.value, expr)
		return result
	}

	for _, item := range matches {
		result = append(result, makeMatch(item, ep.SubexpNames()))
	}

	return result
}

func (sg *String) NotMatch(expr string) interfaces.IString {
	ep, err := regexp.Compile(expr)
	if err != nil {
		sg.LogError("String NotMatch convert error:%v , expr:%v", sg.value, expr)
		return sg
	}

	if ep.MatchString(sg.value) {
		sg.LogError("String NotMatch convert error:%v ", sg.value)
		return sg
	}

	return sg
}
