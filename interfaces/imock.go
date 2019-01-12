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

type IMock interface {
	IReport
	Request(method, path string, args ...interface{}) IRequest
	Options(path string, args ...interface{}) IRequest
	Head(path string, args ...interface{}) IRequest
	Get(path string, args ...interface{}) IRequest
	Post(path string, args ...interface{}) IRequest
	Put(path string, args ...interface{}) IRequest
	Patch(path string, args ...interface{}) IRequest
	Delete(path string, args ...interface{}) IRequest
	Value(value interface{}) IValue
	Object(value map[string]interface{}) IObject
	Array(value []interface{}) IArray
	String(value string) IString
	Float(value float64) IFloat
	Integer(value int) IInteger
	Bool(value bool) IBool
}
