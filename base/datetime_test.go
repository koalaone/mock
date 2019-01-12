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
	"testing"
	"time"

	"github.com/koalaone/mock/reporter"
)

func TestNewDatetime(t *testing.T) {
	newValue := time.Now()

	newDatetime := NewDatetime(newValue)
	if newDatetime.Raw().UnixNano() != newValue.UnixNano() {
		t.Errorf("TestNewDatetime error")
		return
	}
}

func TestDatetime_Raw(t *testing.T) {
	newValue := time.Now()

	newDatetime := NewDatetime(newValue)
	if newDatetime.Raw() != newValue {
		t.Errorf("TestNewDatetime error")
		return
	}
}

func TestDatetime_Equal(t *testing.T) {
	newValue := time.Now()
	checkValue := newValue.UnixNano()

	newDatetime := NewDatetime(newValue)
	newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.Equal(time.Unix(0, checkValue))
}

func TestDatetime_Equal2(t *testing.T) {
	newValue := time.Now()
	checkValue := newValue.Unix()

	newDatetime := NewDatetime(newValue)
	//newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.Equal(time.Unix(checkValue, 0))
}

func TestDatetime_NotEqual(t *testing.T) {
	newValue := time.Now()
	checkValue := newValue.Unix()

	newDatetime := NewDatetime(newValue)
	newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.NotEqual(time.Unix(checkValue, 0))
}

func TestDatetime_NotEqual2(t *testing.T) {
	newValue := time.Now()
	checkValue := newValue.UnixNano()

	newDatetime := NewDatetime(newValue)
	//newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.NotEqual(time.Unix(0, checkValue))
}

func TestDatetime_Gt(t *testing.T) {
	newValue := time.Now()

	newDatetime := NewDatetime(newValue.Add(time.Second))
	newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.Gt(newValue)
}

func TestDatetime_Gt2(t *testing.T) {
	newValue := time.Now()

	newDatetime := NewDatetime(newValue.Add(-time.Second))
	//newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.Gt(newValue)
}

func TestDatetime_Ge(t *testing.T) {
	newValue := time.Now()

	newDatetime := NewDatetime(newValue.Add(time.Second))
	newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.Ge(newValue)

	newDatetime2 := NewDatetime(newValue)
	newDatetime2.Ge(newValue)
}

func TestDatetime_Ge2(t *testing.T) {
	newValue := time.Now()

	newDatetime := NewDatetime(newValue.Add(-time.Second))
	//newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.Ge(newValue)
}

func TestDatetime_Lt(t *testing.T) {
	newValue := time.Now()

	newDatetime := NewDatetime(newValue.Add(-time.Second))
	newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.Lt(newValue)
}

func TestDatetime_Le(t *testing.T) {
	newValue := time.Now()

	newDatetime := NewDatetime(newValue.Add(-time.Second))
	newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.Le(newValue)

	newDatetime2 := NewDatetime(newValue)
	newDatetime2.Le(newValue)
}

func TestDatetime_Le2(t *testing.T) {
	newValue := time.Now()

	newDatetime := NewDatetime(newValue.Add(time.Second))
	//newDatetime.SetReport(reporter.NewTestingReporter(t))
	newDatetime.Le(newValue)
}

func TestDatetime_Between(t *testing.T) {
	newMinValue := time.Now().Add(-time.Minute)
	newMaxValue := time.Now().Add(time.Second)

	newDatetime := NewDatetime(time.Now())
	newDatetime.Between(newMinValue, newMaxValue)
}

func TestDatetime_Between2(t *testing.T) {
	newMinValue := time.Now().Add(-time.Minute)
	newMaxValue := time.Now().Add(-time.Second)

	newDatetime := NewDatetime(time.Now())
	newDatetime.Between(newMinValue, newMaxValue)
}

func TestDatetime_Between3(t *testing.T) {
	newMinValue := time.Now().Add(time.Minute)
	newMaxValue := time.Now().Add(2 * time.Minute)

	newDatetime := NewDatetime(time.Now())
	newDatetime.Between(newMinValue, newMaxValue)
}
