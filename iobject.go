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

package mock

type IObject interface {
	Raw() map[string]interface{}
	Path(path string) IValue
	Schema(value interface{}) IObject
	Equal(value interface{}) IObject
	NotEqual(value interface{}) IObject
	Null() IObject
	NotNull() IObject
	EqualEmpty() IObject
	NotEqualEmpty() IObject
	EqualKeyValue(key string, value interface{}) IObject
	NotEqualKeyValue(key string, value interface{}) IObject
	ContainKey(key string) IObject
	NotContainKey(key string) IObject
	ContainMap(value map[string]interface{}) IObject
	NotContainMap(value map[string]interface{}) IObject
	Keys() []IArray
	Values() []IArray
	Value(key string) IValue
}
