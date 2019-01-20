package httpde_base

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/ajg/form"
	"github.com/koalaone/mock/base"
	"github.com/koalaone/mock/interfaces"
)

type Response struct {
	base.Report
	response *http.Response
	time     time.Duration
}

func NewResponse(resp *http.Response, duration time.Duration) *Response {
	return &Response{
		response: resp,
		time:     duration,
	}
}

func (resp *Response) GetRaw() *http.Response {
	return resp.response
}

func (resp *Response) GetUseTime() interfaces.IInteger {
	return base.NewInteger(int(resp.time.Nanoseconds()))
}

func (resp *Response) GetStatus() interfaces.IInteger {
	return base.NewInteger(resp.response.StatusCode)
}

func (resp *Response) GetStatusPart() interfaces.IInteger {
	switch {
	case resp.response.StatusCode >= interfaces.Status1xx && resp.response.StatusCode < interfaces.Status2xx:
		return base.NewInteger(interfaces.Status1xx)

	case resp.response.StatusCode >= interfaces.Status2xx && resp.response.StatusCode < interfaces.Status3xx:
		return base.NewInteger(interfaces.Status2xx)

	case resp.response.StatusCode >= interfaces.Status3xx && resp.response.StatusCode < interfaces.Status4xx:
		return base.NewInteger(interfaces.Status3xx)

	case resp.response.StatusCode >= interfaces.Status4xx && resp.response.StatusCode < interfaces.Status5xx:
		return base.NewInteger(interfaces.Status4xx)

	case resp.response.StatusCode >= interfaces.Status5xx && resp.response.StatusCode < interfaces.Status6xx:
		return base.NewInteger(interfaces.Status5xx)

	default:
		return base.NewInteger(interfaces.Status6xx)
	}

}

func (resp *Response) GetHeaders() interfaces.IObject {
	headers := make(map[string]interface{})

	if resp.response == nil {
		resp.LogError("GetHeaders response nil")
		return base.NewObject(nil)
	}

	for key, val := range resp.response.Header {
		headers[key] = val
	}

	return base.NewObject(headers)
}

func (resp *Response) GetHeader(key string) interfaces.IString {

	if resp.response == nil {
		resp.LogError("GetHeader response nil")
		return base.NewString("")
	}

	value := resp.response.Header.Get(key)

	return base.NewString(value)
}

func (resp *Response) GetCookies() interfaces.IArray {

	if resp.response == nil {
		resp.LogError("GetCookies response nil")
		return base.NewArray(nil)
	}

	cookies := make([]interface{}, 0)

	for _, val := range resp.response.Cookies() {
		cookies = append(cookies, val)
	}

	return base.NewArray(cookies)
}

func (resp *Response) GetCookie(key string) interfaces.ICookie {

	if resp.response == nil {
		resp.LogError("GetCookie response nil")
		return NewCookie(nil)
	}

	for _, val := range resp.response.Cookies() {
		if val.Name == key {
			return NewCookie(val)
		}
	}

	return NewCookie(nil)
}

func (resp *Response) GetBody() interfaces.IString {
	if resp.response == nil {
		resp.LogError("GetBody response nil")
		return base.NewString("")
	}

	content, err := ioutil.ReadAll(resp.response.Body)
	if err != nil {
		resp.LogError("GetBody ReadAll error:%v", err.Error())
		return base.NewString("")
	}

	return base.NewString(string(content))
}

func (resp *Response) checkContentType(contentType string, charset string) bool {
	locContentType := resp.response.Header.Get("Content-Type")
	if contentType == "" && charset == "" && locContentType == "" {
		return true
	}

	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		resp.LogError("checkContentType ParseMediaType error:%v", err.Error())
		return false
	}

	if mediaType != contentType {
		return false
	}

	locCharset := params["charset"]

	if charset == "" {
		if locCharset != "" && !strings.EqualFold(locCharset, "utf-8") {
			return false
		}
	} else {
		if !strings.EqualFold(locCharset, charset) {
			return false
		}
	}

	return true
}

func (resp *Response) GetBodyText() interfaces.IString {

	if resp.response == nil {
		resp.LogError("GetBodyText response nil")
		return base.NewString("")
	}

	content, err := ioutil.ReadAll(resp.response.Body)
	if err != nil {
		resp.LogError("GetBody ReadAll error:%v", err.Error())
		return base.NewString("")
	}

	if resp.checkContentType("text/plain", "") {
		return base.NewString(string(content))
	}

	return base.NewString("")
}

func (resp *Response) GetBodyForm() interfaces.IObject {

	if resp.response == nil {
		resp.LogError("GetBodyForm response nil")
		return base.NewObject(nil)
	}

	if resp.checkContentType("application/x-www-form-urlencoded", "") {
		content, err := ioutil.ReadAll(resp.response.Body)
		if err != nil {
			resp.LogError("GetBodyForm ReadAll error:%v", err.Error())
			return base.NewObject(nil)
		}

		decoder := form.NewDecoder(bytes.NewReader(content))
		result := make(map[string]interface{})
		err = decoder.Decode(&result)
		if err != nil {
			resp.LogError("GetBodyForm Decode error:%v", err.Error())
			return base.NewObject(nil)
		}

		return base.NewObject(result)
	}

	return base.NewObject(nil)
}

func (resp *Response) GetBodyJSON() interfaces.IObject {

	if resp.response == nil {
		resp.LogError("GetBodyJSON response nil")
		return base.NewObject(nil)
	}

	if resp.checkContentType("application/json", "") {
		content, err := ioutil.ReadAll(resp.response.Body)
		if err != nil {
			resp.LogError("GetBodyJSON ReadAll error:%v", err.Error())
			return base.NewObject(nil)
		}

		var result map[string]interface{}
		err = json.Unmarshal(content, &result)
		if err != nil {
			resp.LogError("GetBodyJSON Decode error:%v", err.Error())
			return base.NewObject(nil)
		}

		return base.NewObject(result)
	}

	return base.NewObject(nil)
}

func (resp *Response) GetBodyJSONP(callback string) interfaces.IValue {
	if resp.response == nil {
		resp.LogError("GetBodyJSONP response nil")
		return base.NewValue(nil)
	}

	if resp.checkContentType("application/javascript", "") {

		content, err := ioutil.ReadAll(resp.response.Body)
		if err != nil {
			resp.LogError("GetBodyJSONP ReadAll error:%v", err.Error())
			return base.NewValue(nil)
		}

		jpExp := regexp.MustCompile(`^\s*([^\s(]+)\s*\((.*)\)\s*;*\s*$`)
		jpResult := jpExp.FindSubmatch(content)
		if len(jpResult) != 3 || string(jpResult[1]) != callback {
			return base.NewValue(nil)
		}

		var result interface{}
		err = json.Unmarshal(jpResult[2], &result)
		if err != nil {
			resp.LogError("GetBodyJSONP Unmarshal error:%v", err.Error())
			return base.NewValue(nil)
		}

		return base.NewValue(result)
	}

	return base.NewValue(nil)
}

func (resp *Response) EqualNoContent() interfaces.IResponse {

	contentType := resp.response.Header.Get("Content-Type")
	if contentType != "" {
		resp.LogError("EqualNoContent Content-Type value is not empty")
		return resp
	}

	content, err := ioutil.ReadAll(resp.response.Body)
	if err != nil {
		resp.LogError("EqualNoContent ReadAll error:%v", err.Error())
		return resp
	}
	if len(content) != 0 {
		resp.LogError("EqualNoContent Body value is not empty")
		return resp
	}

	return resp
}

func (resp *Response) EqualContentType(contentType, characterSet string) interfaces.IResponse {
	if !resp.checkContentType(contentType, characterSet) {
		resp.LogError("EqualContentType equal diff")
		return resp
	}

	return resp
}

func (resp *Response) EqualContentEncoding(contentEncoding string) interfaces.IResponse {

	locContentEncoding := resp.GetHeader("Content-Encoding")
	if locContentEncoding.Raw() != contentEncoding {
		resp.LogError("EqualContentType equal diff")
		return resp
	}

	return resp
}

func (resp *Response) EqualTransferEncoding(transferEncodings ...string) interfaces.IResponse {

	if resp.response == nil {
		resp.LogError("EqualTransferEncoding response nil")
		return resp
	}

	locTransferEncodings := resp.response.TransferEncoding

	if !reflect.DeepEqual(locTransferEncodings, transferEncodings) {
		resp.LogError("EqualTransferEncoding equal diff")
		return resp
	}

	return nil
}
