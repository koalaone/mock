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

type Integer struct {
	Report
	value int
}

func NewInteger(value int) *Integer {
	return &Integer{value: value}
}

func (ir *Integer) Raw() int {
	return ir.value
}

func (ir *Integer) Path(path string) interfaces.IValue {
	newValue, err := convertPath(ir.value, path)
	if err != nil {
		ir.LogError("Integer Path error:%v", err)
		return &Value{value: ir.value}
	}

	return newValue
}

func (ir *Integer) Schema(value int) interfaces.IInteger {
	err := checkSchema(ir.value, value)
	if err != nil {
		ir.LogError("Integer Schema error:%v", err)
	}

	return ir
}

func (ir *Integer) Equal(value int) interfaces.IInteger {
	if !reflect.DeepEqual(ir.value, value) {
		ir.LogError("Integer Equal convert error:%v, target:%v", ir.value, value)
	}

	return ir
}

func (ir *Integer) NotEqual(value int) interfaces.IInteger {
	if reflect.DeepEqual(ir.value, value) {
		ir.LogError("Integer NotEqual convert error:%v, target:%v", ir.value, value)
	}

	return ir
}

func (ir *Integer) Gt(value int) interfaces.IInteger {
	if !(ir.value > value) {
		ir.LogError("Integer Gt convert error:%v, target:%v", ir.value, value)
	}

	return ir
}

func (ir *Integer) Ge(value int) interfaces.IInteger {
	if !(ir.value >= value) {
		ir.LogError("Integer Ge convert error:%v, target:%v", ir.value, value)
	}

	return ir
}

func (ir *Integer) Lt(value int) interfaces.IInteger {
	if !(ir.value < value) {
		ir.LogError("Integer Lt convert error:%v, target:%v", ir.value, value)
	}

	return ir
}

func (ir *Integer) Le(value int) interfaces.IInteger {
	if !(ir.value <= value) {
		ir.LogError("Integer Le convert error:%v, target:%v", ir.value, value)
	}

	return ir
}

func (ir *Integer) Between(min, max int) interfaces.IInteger {
	if !(ir.value >= min && ir.value <= max) {
		ir.LogError("Integer Le convert error:%v, min:%v, max:%v", ir.value, min, max)
	}

	return ir
}
