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
	"github.com/koalaone/mock/reporter"
)

type Bool struct {
	Report
	value bool
}

func NewBool(value bool) *Bool {
	report := reporter.NewOSReporter()
	newBool := &Bool{
		value: value,
	}
	newBool.report = report
	return newBool
}

func (bl *Bool) Raw() bool {
	return bl.value
}

func (bl *Bool) Path(path string) interfaces.IValue {
	newValue, err := convertPath(bl.value, path)
	if err != nil {
		bl.LogError("Bool Path error:%v", err)
		return NewValue(bl.value)
	}

	return newValue
}

func (bl *Bool) Schema(value interface{}) interfaces.IBool {
	err := checkSchema(bl.value, value)
	if err != nil {
		bl.LogError("Bool Schema error:%v", err)
	}

	return bl
}

func (bl *Bool) Equal(value interface{}) interfaces.IBool {
	if !reflect.DeepEqual(bl.value, value) {
		bl.LogError("Bool Equal convert error:%v, target:%v", bl.value, value)
	}

	return bl
}

func (bl *Bool) NotEqual(value interface{}) interfaces.IBool {
	if reflect.DeepEqual(bl.value, value) {
		bl.LogError("Bool NotEqual convert error:%v, target:%v", bl.value, value)
	}

	return bl
}

func (bl *Bool) True() interfaces.IBool {
	return bl.Equal(true)
}

func (bl *Bool) False() interfaces.IBool {
	return bl.Equal(false)
}
