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
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/koalaone/mock/interfaces"
	"github.com/xeipuuv/gojsonschema"
	"github.com/yalp/jsonpath"
)

func convertValue(value interface{}) (interface{}, error) {
	obj, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	var out interface{}
	err = json.Unmarshal(obj, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func convertPath(value interface{}, path string) (interfaces.IValue, error) {
	value, err := jsonpath.Read(value, path)
	if err != nil {
		return nil, err
	}

	return &Value{value: value}, nil
}

func convertString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func checkSchema(value, schema interface{}) error {
	sourceLoader := gojsonschema.NewGoLoader(value)
	var schemaLoader gojsonschema.JSONLoader

	schemaSource := convertString(schema)
	ok, err := regexp.MatchString(`^\w+://`, schemaSource)
	if err != nil {
		return err
	}
	if ok {
		schemaLoader = gojsonschema.NewReferenceLoader(schemaSource)
	} else {
		schemaLoader = gojsonschema.NewStringLoader(schemaSource)
	}

	result, err := gojsonschema.Validate(schemaLoader, sourceLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		errorList := ""
		for _, err := range result.Errors() {
			errorList += fmt.Sprintf(" %s\n", err)
		}

		return errors.New(errorList)
	}

	return nil
}
