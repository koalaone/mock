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
	"time"

	"github.com/koalaone/mock/interfaces"
	"github.com/koalaone/mock/reporter"
)

type Datetime struct {
	Report
	value time.Time
}

func NewDatetime(value time.Time) interfaces.IDatetime {
	report := reporter.NewOSReporter()
	newDatetime := &Datetime{
		value: value,
	}
	newDatetime.report = report
	return newDatetime
}

func (dt *Datetime) Raw() time.Time {
	return dt.value
}

func (dt *Datetime) Equal(value time.Time) interfaces.IDatetime {
	if !dt.value.Equal(value) {
		dt.LogError("Datetime Equal convert error:%v, target:%v", dt.value, value)
	}

	return dt
}

func (dt *Datetime) NotEqual(value time.Time) interfaces.IDatetime {
	if dt.value.Equal(value) {
		dt.LogError("Datetime NotEqual convert error:%v, target:%v", dt.value, value)
	}

	return dt
}

func (dt *Datetime) Gt(value time.Time) interfaces.IDatetime {
	if !dt.value.After(value) {
		dt.LogError("Datetime Gt convert error:%v, target:%v", dt.value, value)
	}

	return dt
}

func (dt *Datetime) Ge(value time.Time) interfaces.IDatetime {
	if !(dt.value.After(value) || dt.value.Equal(value)) {
		dt.LogError("Datetime Ge convert error:%v, target:%v", dt.value, value)
	}

	return dt
}

func (dt *Datetime) Lt(value time.Time) interfaces.IDatetime {
	if !dt.value.Before(value) {
		dt.LogError("Datetime Lt convert error:%v, target:%v", dt.value, value)
	}

	return dt
}

func (dt *Datetime) Le(value time.Time) interfaces.IDatetime {
	if !(dt.value.Before(value) || dt.value.Equal(value)) {
		dt.LogError("Datetime Le convert error:%v, target:%v", dt.value, value)
	}

	return dt
}

func (dt *Datetime) Between(min, max time.Time) interfaces.IDatetime {
	if !((dt.value.After(min) || dt.value.Equal(min)) &&
		(dt.value.Before(max) || dt.value.Equal(max))) {
		dt.LogError("Datetime Between convert error:%v, between min:%v, max:%v", min, max)
	}

	return dt
}
