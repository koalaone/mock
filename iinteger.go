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

type IInteger interface {
	Raw() int
	Path(path string) IValue
	Schema(value interface{}) IInteger
	Equal(value interface{}) IInteger
	NotEqual(value interface{}) IInteger
	Null() IInteger
	NotNull() IInteger
	Gt(value interface{}) IInteger
	Ge(value interface{}) IInteger
	Lt(value interface{}) IInteger
	Le(value interface{}) IInteger
	Between(min, max interface{}) IInteger
}
