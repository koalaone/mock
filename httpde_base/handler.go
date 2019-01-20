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
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/koalaone/mock/base"
)

type Handler struct {
	base.Report
	Handler http.Handler
	TLS     *tls.ConnectionState
}

func NewHandler(handler http.Handler) Handler {
	return Handler{Handler: handler}
}

func (hr *Handler) SetTLS(value *tls.ConnectionState) *Handler {
	if value == nil {
		return hr
	}

	hr.TLS = value

	return hr
}

func (hr Handler) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Proto == "" {
		req.Proto = fmt.Sprintf("HTTP/%d.%d", req.ProtoMajor, req.ProtoMinor)
	}

	if req.Body != nil {
		if req.ContentLength == -1 {
			req.TransferEncoding = []string{"chunked"}
		}
	} else {
		req.Body = ioutil.NopCloser(bytes.NewReader(nil))
	}

	if req.URL != nil && req.URL.Scheme == "https" && hr.TLS != nil {
		req.TLS = hr.TLS
	}

	if req.RequestURI == "" {
		req.RequestURI = req.URL.RequestURI()
	}

	recorder := httptest.NewRecorder()

	hr.Handler.ServeHTTP(recorder, req)

	resp := http.Response{
		Request:    req,
		StatusCode: recorder.Code,
		Status:     http.StatusText(recorder.Code),
		Header:     recorder.HeaderMap,
	}

	if recorder.Flushed {
		resp.TransferEncoding = []string{"chunked"}
	}

	if recorder.Body != nil {
		resp.Body = ioutil.NopCloser(recorder.Body)
	}

	return &resp, nil
}
