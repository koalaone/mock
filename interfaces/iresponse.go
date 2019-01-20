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

package interfaces

import (
	"net/http"
)

const (
	Status1xx = 100
	Status2xx = 200
	Status3xx = 300
	Status4xx = 400
	Status5xx = 500
	Status6xx = 600
)

type IResponse interface {
	IReport
	GetRaw() *http.Response
	GetUseTime() IInteger
	GetStatus() IInteger
	GetStatusPart() IInteger
	GetHeaders() IObject
	GetHeader(key string) IString
	GetCookies() IArray
	GetCookie(key string) ICookie
	GetBody() IString
	GetBodyText() IString
	GetBodyForm() IObject
	GetBodyJSON() IObject
	GetBodyJSONP(callback string) IValue
	EqualNoContent() IResponse
	EqualContentType(contentType, characterSet string) IResponse
	EqualContentEncoding(contentEncoding string) IResponse
	EqualTransferEncoding(transferEncodings ...string) IResponse
}
