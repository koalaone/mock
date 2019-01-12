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
	"math"
	"reflect"

	"github.com/koalaone/mock/interfaces"
)

type Float struct {
	Report
	value float64
}

func NewFloat(value float64) *Float {
	return &Float{value: value}
}

func (ft *Float) Raw() float64 {
	return ft.value
}

func (ft *Float) Path(path string) interfaces.IValue {
	newValue, err := convertPath(ft.value, path)
	if err != nil {
		ft.LogError("Float Path error:%v", err)
		return &Value{value: ft.value}
	}

	return newValue
}

func (ft *Float) Schema(value float64) interfaces.IFloat {
	err := checkSchema(ft.value, value)
	if err != nil {
		ft.LogError("Float Schema error:%v", err)
	}

	return ft
}

func (ft *Float) Equal(value float64) interfaces.IFloat {
	if !reflect.DeepEqual(ft.value, value) {
		ft.LogError("Float Equal convert error:%v, target:%v", ft.value, value)
	}

	return ft
}

func (ft *Float) NotEqual(value float64) interfaces.IFloat {
	if !reflect.DeepEqual(ft.value, value) {
		ft.LogError("Float NotEqual convert error:%v, target:%v", ft.value, value)
	}

	return ft
}

func (ft *Float) EqualDeviation(value, deviation float64) interfaces.IFloat {
	if math.IsNaN(ft.value) || math.IsNaN(value) || math.IsNaN(deviation) {
		ft.LogError("Float EqualDeviation convert error:%v, value:%v, deviation:%v", ft.value, value, deviation)
		return ft
	}

	diffValue := ft.value - value
	if diffValue < -deviation || diffValue > deviation {
		ft.LogError("Float EqualDeviation convert error:%v, value:%v, deviation:%v", ft.value, value, deviation)
		return ft
	}

	return ft
}

func (ft *Float) NotEqualDeviation(value, deviation float64) interfaces.IFloat {
	if math.IsNaN(ft.value) || math.IsNaN(value) || math.IsNaN(deviation) {
		ft.LogError("Float NotEqualDeviation convert error:%v, value:%v, deviation:%v", ft.value, value, deviation)
		return ft
	}

	diffValue := ft.value - value
	if !(diffValue < -deviation || diffValue > deviation) {
		ft.LogError("Float NotEqualDeviation convert error:%v, value:%v, deviation:%v", ft.value, value, deviation)
		return ft
	}

	return ft
}

func (ft *Float) Gt(value float64) interfaces.IFloat {
	if !(ft.value > value) {
		ft.LogError("Float Gt convert error:%v, target:%v", ft.value, value)
	}

	return ft
}

func (ft *Float) Ge(value float64) interfaces.IFloat {
	if !(ft.value >= value) {
		ft.LogError("Float Ge convert error:%v, target:%v", ft.value, value)
	}

	return ft
}

func (ft *Float) Lt(value float64) interfaces.IFloat {
	if !(ft.value < value) {
		ft.LogError("Float Lt convert error:%v, target:%v", ft.value, value)
	}

	return ft
}

func (ft *Float) Le(value float64) interfaces.IFloat {
	if !(ft.value <= value) {
		ft.LogError("Float Le convert error:%v, target:%v", ft.value, value)
	}

	return ft
}

func (ft *Float) Between(min, max float64) interfaces.IFloat {
	if !(ft.value >= min && ft.value <= max) {
		ft.LogError("Float Le convert error:%v, min:%v, max:%v", ft.value, min, max)
	}

	return ft
}
