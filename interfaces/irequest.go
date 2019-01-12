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
	"io"
	"net/http"
)

type IRequest interface {
	IReport
	SetClient(client IClient) IRequest
	SetDomain(value string) IRequest
	SetHandler(handler http.Handler) IRequest
	SetMethod(method string) IRequest
	SetPath(path string) IRequest
	SetPathParam(key string, value interface{}) IRequest
	SetPathParamObject(object interface{}) IRequest
	SetPathQuery(key string, value interface{}) IRequest
	SetPathQueryObject(object interface{}) IRequest
	SetPathQueryString(query string) IRequest
	SetHeaders(headers map[string]string) IRequest
	SetHeader(key, value string) IRequest
	SetCookies(cookies map[string]string) IRequest
	SetCookie(key, value string) IRequest
	SetBasicAuth(username, password string) IRequest
	SetHttpVersion(version string) IRequest
	SetBodyMultipart() IRequest
	SetBodyChunked(body io.Reader) IRequest
	SetBodyBytes(body []byte) IRequest
	SetBodyText(body string) IRequest
	SetBodyJson(body interface{}) IRequest
	SetBodyForm(body interface{}) IRequest
	SetBodyFormField(key string, value interface{}) IRequest
	SetBodyFile(key, filename string, reader io.Reader) IRequest
	SetBodyFileBytes(key, filename string, data []byte) IRequest
	Mock() IResponse
}
