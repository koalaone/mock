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

type Array struct {
	Report
	value []interface{}
}

func NewArray(value []interface{}) *Array {
	newArray := &Array{
		value: make([]interface{}, 0),
	}
	if value != nil {
		newArray.value = append(newArray.value, value)
	}

	return newArray
}

func (ay *Array) Raw() []interface{} {
	return ay.value
}

func (ay *Array) Path(path string) interfaces.IValue {
	newValue, err := convertPath(ay.value, path)
	if err != nil {
		ay.LogError("Array Path error:%v", err)
		return &Value{value: ay.value}
	}

	return newValue
}

func (ay *Array) Schema(value interface{}) interfaces.IArray {
	err := checkSchema(ay.value, value)
	if err != nil {
		ay.LogError("Array Schema error:%v", err)
	}

	return ay
}

func (ay *Array) Equal(value interface{}) interfaces.IArray {
	if !reflect.DeepEqual(ay.value, value) {
		ay.LogError("Array Equal convert error:%v, target:%v", ay.value, value)
	}

	return ay
}

func (ay *Array) NotEqual(value interface{}) interfaces.IArray {
	if reflect.DeepEqual(ay.value, value) {
		ay.LogError("Array NotEqual convert error:%v, target:%v", ay.value, value)
	}

	return ay
}

func (ay *Array) Null() interfaces.IArray {
	if ay.value != nil {
		ay.LogError("Array Null convert error")
	}

	return ay
}

func (ay *Array) NotNull() interfaces.IArray {
	if ay.value != nil {
		ay.LogError("Array NotNull convert error")
	}

	return ay
}

func (ay *Array) EqualEmpty() interfaces.IArray {
	return ay.Equal([]interface{}{})
}

func (ay *Array) NotEqualEmpty() interfaces.IArray {
	return ay.NotEqual([]interface{}{})
}

func (ay *Array) EqualElements(values ...interface{}) interfaces.IArray {
	return ay.Equal(values)
}

func (ay *Array) containsElement(value interface{}) bool {
	for _, item := range ay.value {
		if reflect.DeepEqual(value, item) {
			return true
		}
	}

	return false
}

func (ay *Array) Contains(values ...interface{}) interfaces.IArray {
	for _, item := range values {
		if !ay.containsElement(item) {
			ay.LogError("Array Contains check error:%v, target:%v", ay.value, values)
		}
	}

	return ay
}

func (ay *Array) NotContains(values ...interface{}) interfaces.IArray {
	for _, item := range values {
		if ay.containsElement(item) {
			ay.LogError("Array NotContains check error:%v, target:%v", ay.value, values)
		}
	}

	return ay
}

func (ay *Array) Length() interfaces.IInteger {
	return NewInteger(len(ay.value))
}

func (ay *Array) Get(index int) interfaces.IValue {
	if index < 0 || index >= len(ay.value) {
		ay.LogError("Array Get check error: out of range")
		return NewValue(ay.value)
	}

	return NewValue(ay.value[index])
}

func (ay *Array) First() interfaces.IValue {
	if len(ay.value) <= 0 {
		ay.LogError("Array First check error: out of range")
		return NewValue(ay.value)
	}

	return NewValue(ay.value[0])
}

func (ay *Array) Last() interfaces.IValue {
	if len(ay.value) <= 0 {
		ay.LogError("Array Last check error: out of range")
		return NewValue(ay.value)
	}

	return NewValue(ay.value[len(ay.value)-1])
}

func (ay *Array) Iterator() []interfaces.IValue {
	retList := make([]interfaces.IValue, 0)
	for n := range ay.value {
		retList = append(retList, NewValue(ay.value[n]))
	}

	return retList
}
