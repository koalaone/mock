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

package http

import (
	"net/http"

	"github.com/koalaone/mock/base"
	"github.com/koalaone/mock/interfaces"
)

type Cookie struct {
	base.Report
	value *http.Cookie
}

func NewCookie(value *http.Cookie) *Cookie {
	newCookie := &Cookie{
		value: value,
	}

	return newCookie
}

func (ce *Cookie) Raw() *http.Cookie {
	return ce.value
}

func (ce *Cookie) Name() interfaces.IString {
	return base.NewString(ce.value.Name)
}

func (ce *Cookie) Value() interfaces.IString {
	return base.NewString(ce.value.Value)
}

func (ce *Cookie) Domain() interfaces.IString {
	return base.NewString(ce.value.Domain)
}

func (ce *Cookie) Path() interfaces.IString {
	return base.NewString(ce.value.Path)
}

func (ce *Cookie) Expires() interfaces.IDatetime {
	return base.NewDatetime(ce.value.Expires)
}
