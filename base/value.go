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

type Value struct {
	Report
	value interface{}
}

func NewValue(value interface{}) *Value {
	newValue, err := convertValue(value)
	if err != nil {
		return &Value{value: nil}
	}

	return &Value{value: newValue}
}

func (ve *Value) Raw() interface{} {
	return ve.value
}

func (ve *Value) Path(path string) interfaces.IValue {
	newValue, err := convertPath(ve.value, path)
	if err != nil {
		ve.LogError("Value Path error:%v", err)
		return ve
	}

	return newValue
}

func (ve *Value) Schema(value interface{}) interfaces.IValue {
	err := checkSchema(ve.value, value)
	if err != nil {
		ve.LogError("Value Schema error:%v", err)
	}

	return ve
}

func (ve *Value) Object() interfaces.IObject {
	cObj, ok := ve.value.(map[string]interface{})
	if !ok {
		ve.LogError("Value Object convert error:%v", ve.value)

		return NewObject(make(map[string]interface{}, 0))
	}

	return NewObject(cObj)
}

func (ve *Value) Array() interfaces.IArray {
	cArray, ok := ve.value.([]interface{})
	if !ok {
		ve.LogError("Value Array convert error:%v", ve.value)

		return NewArray(make([]interface{}, 0))
	}

	return NewArray(cArray)
}

func (ve *Value) String() interfaces.IString {
	value, ok := ve.value.(string)
	if !ok {
		ve.LogError("Value String convert error")

		return NewString("")
	}

	return NewString(value)
}

func (ve *Value) Integer() interfaces.IInteger {
	value, ok := ve.value.(int)
	if !ok {
		ve.LogError("Value Integer convert error")

		return NewInteger(0)
	}

	return NewInteger(value)
}

func (ve *Value) Float() interfaces.IFloat {
	value, ok := ve.value.(float64)
	if !ok {
		ve.LogError("Value Float convert error")

		return NewFloat(0)
	}

	return NewFloat(value)
}

func (ve *Value) Bool() interfaces.IBool {
	value, ok := ve.value.(bool)
	if !ok {
		ve.LogError("Value Bool convert error")

		return NewBool(false)
	}

	return NewBool(value)
}

func (ve *Value) Null() interfaces.IValue {
	if ve.value != nil {
		ve.LogError("Value Null convert error:%v", ve.value)
	}

	return ve
}

func (ve *Value) NotNull() interfaces.IValue {
	if ve.value == nil {
		ve.LogError("Value NotNull convert error:%v", ve.value)
	}

	return ve
}

func (ve *Value) Equal(value interface{}) interfaces.IValue {
	if !reflect.DeepEqual(ve.value, value) {
		ve.LogError("Value Equal convert error:%v, target:%v", ve.value, value)
	}

	return ve
}

func (ve *Value) NotEqual(value interface{}) interfaces.IValue {
	if reflect.DeepEqual(ve.value, value) {
		ve.LogError("Value NotEqual convert error:%v, target:%v", ve.value, value)
	}

	return ve
}
