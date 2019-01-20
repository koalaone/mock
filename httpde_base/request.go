package httpde_base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/koalaone/mock/base"
	"github.com/koalaone/mock/interfaces"
)

type Request struct {
	base.Report
	client         interfaces.IClient
	request        *http.Request
	pathArgs       []interface{}
	query          url.Values
	form           url.Values
	formBuf        *bytes.Buffer
	multipart      *multipart.Writer
	setContentType bool
}

func NewRequest(method, path string, pathArgs ...interface{}) *Request {
	var err error

	newRequest := &Request{}
	newRequest.client = NewHttpClient()
	newRequest.pathArgs = append(newRequest.pathArgs, pathArgs...)
	newRequest.request, err = http.NewRequest(method, path, nil)
	if err != nil {
		newRequest.LogError("NewRequest error:%v", err.Error())
	}

	return newRequest
}

func (req *Request) SetClient(client interfaces.IClient) interfaces.IRequest {
	if client == nil {
		return req
	}

	req.client = client

	return req
}

func (req *Request) SetHandler(handler http.Handler) interfaces.IRequest {
	if handler == nil {
		return req
	}

	if req.client == nil {
		req.client = NewHttpClient()
	}

	req.client.Client().Transport = NewHandler(handler)

	return req
}

func (req *Request) SetPathQuery(key string, value interface{}) interfaces.IRequest {
	if key == "" {
		return req
	}

	if req.query == nil {
		req.query = make(url.Values, 0)
	}

	req.query.Add(key, fmt.Sprint(value))

	return req
}

func (req *Request) SetPathQueryObject(object interface{}) interfaces.IRequest {
	if object == nil {
		return nil
	}

	beforeValue, err := json.Marshal(object)
	if err != nil {
		return req
	}

	midResultMap := make(map[string]interface{}, 0)
	jen := json.NewDecoder(bytes.NewBuffer(beforeValue))
	jen.UseNumber()
	err = jen.Decode(&midResultMap)
	if err != nil {
		return req
	}

	for key, val := range midResultMap {
		req.SetPathQuery(key, val)
	}

	return req
}

func (req *Request) SetPathQueryString(query string) interfaces.IRequest {
	if query == "" {
		return req
	}

	querys, err := url.ParseQuery(query)
	if err != nil {
		return req
	}

	for key, val := range querys {
		req.SetPathQuery(key, val)
	}

	return req
}

func (req *Request) SetHeaders(headers map[string]string) interfaces.IRequest {
	if headers == nil {
		return req
	}

	for key, val := range headers {
		req.request.Header.Add(key, val)
	}

	return req
}

func (req *Request) SetHeader(key, value string) interfaces.IRequest {
	if key == "" {
		return req
	}

	switch http.CanonicalHeaderKey(key) {
	case "Host":
		req.request.Host = value
	case "Content-Type":
		if !req.setContentType {
			delete(req.request.Header, "Content-Type")
		}
		req.setContentType = true
		req.request.Header.Add(key, value)
	default:
		req.request.Header.Add(key, value)
	}

	return req
}

func (req *Request) SetCookies(cookies map[string]string) interfaces.IRequest {
	if cookies == nil {
		return req
	}

	for key, val := range cookies {
		req.SetCookie(key, val)
	}

	return req
}

func (req *Request) SetCookie(key, value string) interfaces.IRequest {
	if key == "" {
		return req
	}

	req.request.AddCookie(&http.Cookie{Name: key, Value: value})

	return req
}

func (req *Request) SetBasicAuth(username, password string) interfaces.IRequest {

	req.request.SetBasicAuth(username, password)

	return req
}

func (req *Request) SetHttpVersion(version string) interfaces.IRequest {

	major, minor, ok := http.ParseHTTPVersion(version)
	if !ok {
		return req
	}

	req.request.ProtoMajor = major
	req.request.ProtoMinor = minor

	return req
}

func (req *Request) SetBodyMultipart() interfaces.IRequest {

	req.SetHeader("Content-Type", "multipart/form-data")
	req.multipart = multipart.NewWriter(req.formBuf)

	if req.formBuf == nil {
		req.request.Body = nil
		req.request.ContentLength = 0
	} else {
		req.request.Body = ioutil.NopCloser(req.formBuf)
		req.request.ContentLength = int64(req.formBuf.Len())
	}

	return req
}

func (req *Request) SetBodyChunked(body io.Reader) interfaces.IRequest {
	req.SetHeader("Content-Type", "application/octet-stream")

	if body == nil {
		req.request.Body = nil
		req.request.ContentLength = 0
	} else {
		req.request.Body = ioutil.NopCloser(body)
		req.request.ContentLength = -1
	}

	return req
}

func (req *Request) SetBodyBytes(body []byte) interfaces.IRequest {
	if body == nil {
		req.request.Body = nil
		req.request.ContentLength = 0
	} else {
		req.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		req.request.ContentLength = int64(len(body))
	}

	return req
}

func (req *Request) SetBodyText(body string) interfaces.IRequest {
	req.SetHeader("Content-Type", "text/plain; charset=utf-8")

	if body == "" {
		req.request.Body = nil
		req.request.ContentLength = 0
	} else {
		req.request.Body = ioutil.NopCloser(strings.NewReader(body))
		req.request.ContentLength = int64(len(body))
	}

	return req
}

func (req *Request) SetBodyJson(body interface{}) interfaces.IRequest {
	req.SetHeader("Content-Type", "application/json; charset=utf-8")

	valueBody, err := json.Marshal(body)
	if err != nil {
		req.LogError("SetBodyJson json.Marshal error:%v", err.Error())
		return req
	}
	req.request.Body = ioutil.NopCloser(bytes.NewBuffer(valueBody))
	req.request.ContentLength = int64(len(valueBody))

	return req
}

func (req *Request) SetBodyForm(values map[string]interface{}) interfaces.IRequest {

	for key, val := range values {
		req.SetBodyFormField(key, val)
	}

	return req
}

func (req *Request) SetBodyFormField(key string, value interface{}) interfaces.IRequest {
	if req.multipart == nil {
		req.SetHeader("Content-Type", "application/x-www-form-urlencoded")

		if req.form == nil {
			req.form = make(url.Values)
		}

		locValue := fmt.Sprint(value)
		req.form[key] = append(req.form[key], locValue)

	} else {
		req.SetHeader("Content-Type", "multipart/form-data")

		locValue := fmt.Sprint(value)
		err := req.multipart.WriteField(key, locValue)
		if err != nil {
			req.LogError("SetBodyForm WriteField error:%v", err.Error())
			return req
		}
	}

	return req
}

func (req *Request) SetBodyFile(key, filename string, reader io.Reader) interfaces.IRequest {
	req.SetHeader("Content-Type", "multipart/form-data")

	if req.multipart == nil {
		req.multipart = multipart.NewWriter(req.formBuf)
	}
	locMultipart, err := req.multipart.CreateFormFile(key, filename)
	if err != nil {
		req.LogError("SetBodyFile CreateFormFile error:%v", err.Error())
		return req
	}

	var locReader io.Reader
	if reader != nil {
		locReader = reader
	} else {
		tmp, err := os.Open(filename)
		if err != nil {
			req.LogError("SetBodyFile open file error:%v", err.Error())
			return req
		}
		defer tmp.Close()

		locReader = tmp
	}

	_, err = io.Copy(locMultipart, locReader)
	if err != nil {
		req.LogError("SetBodyFile file Copy error:%v", err.Error())
		return req
	}

	return req
}

func (req *Request) SetBodyFileBytes(key, filename string, data []byte) interfaces.IRequest {

	return req.SetBodyFile(key, filename, bytes.NewReader(data))
}

func (req *Request) makeRequest() {

	if req.query != nil {
		req.request.URL.RawQuery = req.query.Encode()
	}

	if req.multipart != nil {
		err := req.multipart.Close()
		if err != nil {
			req.LogError("Mock multipart Close error:%v", err.Error())

		}

		req.setContentType = false
		req.SetHeader("Content-Type", req.multipart.FormDataContentType())
		req.request.Body = ioutil.NopCloser(req.formBuf)
		req.request.ContentLength = int64(req.formBuf.Len())
	} else {
		if req.form != nil {
			req.setContentType = false
			req.SetHeader("Content-Type", "multipart/form-data")
			for key, val := range req.form {
				req.SetBodyFormField(key, val)
			}
		}
	}

}

func (req *Request) Mock() interfaces.IResponse {

	req.makeRequest()

	start := time.Now()

	resp, err := req.client.Do(req.request)
	if err != nil {
		req.LogError("Mock Do error:%v", err.Error())
		return nil
	}

	useTime := time.Since(start)

	req.LogInfo("Mock use time:%v", useTime.String())

	return NewResponse(resp, useTime)
}
