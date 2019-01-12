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

type Object struct {
	Report
	value map[string]interface{}
}

func NewObject(value map[string]interface{}) *Object {
	if value == nil {
		return &Object{value: make(map[string]interface{}, 0)}
	}

	return &Object{value: value}
}

func (ot *Object) Raw() map[string]interface{} {
	return ot.value
}

func (ot *Object) Path(path string) interfaces.IValue {
	newValue, err := convertPath(ot.value, path)
	if err != nil {
		ot.LogError("Object Path error:%v", err)
		return &Value{value: ot.value}
	}

	return newValue
}

func (ot *Object) Schema(value interface{}) interfaces.IObject {
	err := checkSchema(ot.value, value)
	if err != nil {
		ot.LogError("Object Schema error:%v", err)
	}

	return ot
}

func (ot *Object) Equal(value interface{}) interfaces.IObject {
	if !reflect.DeepEqual(ot.value, value) {
		ot.LogError("Object Equal convert error:%v, target:%v", ot.value, value)
	}

	return ot
}

func (ot *Object) NotEqual(value interface{}) interfaces.IObject {
	if reflect.DeepEqual(ot.value, value) {
		ot.LogError("Object NotEqual convert error:%v, target:%v", ot.value, value)
	}

	return ot
}

func (ot *Object) Null() interfaces.IObject {
	if !(ot.value == nil) {
		ot.LogError("Object Null convert error:%v", ot.value)
	}

	return ot
}

func (ot *Object) NotNull() interfaces.IObject {
	if ot.value == nil {
		ot.LogError("Object NotNull convert error:%v", ot.value)
	}

	return ot
}

func (ot *Object) EqualEmpty() interfaces.IObject {
	return ot.Equal(map[string]interface{}{})
}

func (ot *Object) NotEqualEmpty() interfaces.IObject {
	return ot.NotEqual(map[string]interface{}{})
}

func (ot *Object) EqualKeyValue(key string, value interface{}) interfaces.IObject {
	locValue, ok := ot.value[key]
	if !ok {
		ot.LogError("Object EqualKeyValue convert error:%v", ot.value)
		return ot
	}

	if !reflect.DeepEqual(locValue, value) {
		ot.LogError("Object EqualKeyValue convert error:%v, target:%v", ot.value, value)
	}

	return ot
}

func (ot *Object) NotEqualKeyValue(key string, value interface{}) interfaces.IObject {
	locValue, ok := ot.value[key]
	if !ok {
		ot.LogError("Object NotEqualKeyValue convert error:%v", ot.value)
		return ot
	}

	if reflect.DeepEqual(locValue, value) {
		ot.LogError("Object NotEqualKeyValue convert error:%v, target:%v", ot.value, value)
	}

	return ot
}

func (ot *Object) ContainKey(key string) interfaces.IObject {
	_, ok := ot.value[key]
	if !ok {
		ot.LogError("Object ContainKey convert error:%v", ot.value)
		return ot
	}

	return ot
}

func (ot *Object) NotContainKey(key string) interfaces.IObject {
	_, ok := ot.value[key]
	if ok {
		ot.LogError("Object NotContainKey convert error:%v", ot.value)
		return ot
	}

	return ot
}

func (ot *Object) ContainMap(value map[string]interface{}) interfaces.IObject {
	for key, value := range value {
		sourceValue, sourceOK := ot.value[key]
		if !sourceOK {
			ot.LogError("Object ContainMap convert error:%v, target:%v", ot.value, value)
			return ot
		}
		if !reflect.DeepEqual(sourceValue, value) {
			ot.LogError("Object ContainMap convert error:%v, target:%v", ot.value, value)
			return ot
		}
	}

	return ot
}

func (ot *Object) NotContainMap(value map[string]interface{}) interfaces.IObject {
	for key, value := range value {
		sourceValue, sourceOK := ot.value[key]
		if !sourceOK {
			ot.LogError("Object ContainMap convert error:%v, target:%v", ot.value, value)
			return ot
		}
		if reflect.DeepEqual(sourceValue, value) {
			ot.LogError("Object ContainMap convert error:%v, target:%v", ot.value, value)
			return ot
		}
	}

	return ot
}

func (ot *Object) Keys() []interfaces.IValue {
	result := make([]interfaces.IValue, 0)
	for key := range ot.value {
		result = append(result, NewValue(key))
	}

	return result
}

func (ot *Object) Values() []interfaces.IValue {
	result := make([]interfaces.IValue, 0)
	for _, Item := range ot.value {
		result = append(result, NewValue(Item))
	}

	return result
}

func (ot *Object) Value(key string) interfaces.IValue {
	value, ok := ot.value[key]
	if !ok {
		ot.LogError("Object NotContainKey convert error:%v", ot.value)
		return NewValue(nil)
	}

	return NewValue(value)
}
