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

package httpde_base

import (
	"github.com/koalaone/mock/base"
	"github.com/koalaone/mock/interfaces"
)

type HttpFactory struct {
	base.Report
	Domain   string
	requests map[string]interfaces.IRequest
}

func NewHttpFactory(domain string) *HttpFactory {
	newHttpFactory := HttpFactory{}
	newHttpFactory.Domain = domain

	return &newHttpFactory
}

func (hf *HttpFactory) SetDomain(value string) interfaces.IHttpFactory {
	if value == "" {
		return hf
	}

	hf.Domain = value

	return hf
}

func (hf *HttpFactory) Request(method, path string, pathArgs ...interface{}) interfaces.IRequest {
	req := NewRequest(method, path, pathArgs...)

	if _, ok := hf.requests[path]; !ok {
		hf.requests[path] = req
	}

	return req
}

func (hf *HttpFactory) POST(path string, pathArgs ...interface{}) interfaces.IRequest {
	return hf.Request("POST", path, pathArgs...)
}

func (hf *HttpFactory) GET(path string, pathArgs ...interface{}) interfaces.IRequest {
	return hf.Request("GET", path, pathArgs...)
}

func (hf *HttpFactory) DELETE(path string, pathArgs ...interface{}) interfaces.IRequest {
	return hf.Request("DELETE", path, pathArgs...)
}

func (hf *HttpFactory) PATCH(path string, pathArgs ...interface{}) interfaces.IRequest {
	return hf.Request("PATCH", path, pathArgs...)
}

func (hf *HttpFactory) PUT(path string, pathArgs ...interface{}) interfaces.IRequest {
	return hf.Request("PUT", path, pathArgs...)
}

func (hf *HttpFactory) OPTIONS(path string, pathArgs ...interface{}) interfaces.IRequest {
	return hf.Request("OPTIONS", path, pathArgs...)
}

func (hf *HttpFactory) HEAD(path string, pathArgs ...interface{}) interfaces.IRequest {
	return hf.Request("HEAD", path, pathArgs...)
}

func (hf *HttpFactory) Value(value interface{}) interfaces.IValue {
	return base.NewValue(value)
}

func (hf *HttpFactory) Object(value map[string]interface{}) interfaces.IObject {
	return base.NewObject(value)
}

func (hf *HttpFactory) Array(value []interface{}) interfaces.IArray {
	return base.NewArray(value)
}

func (hf *HttpFactory) String(value string) interfaces.IString {
	return base.NewString(value)
}

func (hf *HttpFactory) Float(value float64) interfaces.IFloat {
	return base.NewFloat(value)
}

func (hf *HttpFactory) Integer(value int) interfaces.IInteger {
	return base.NewInteger(value)
}

func (hf *HttpFactory) Boolean(value bool) interfaces.IBool {
	return base.NewBool(value)
}
