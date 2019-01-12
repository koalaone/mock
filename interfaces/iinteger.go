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

type IInteger interface {
	IReport
	Raw() int
	Path(path string) IValue
	Schema(value int) IInteger
	Equal(value int) IInteger
	NotEqual(value int) IInteger
	Gt(value int) IInteger
	Ge(value int) IInteger
	Lt(value int) IInteger
	Le(value int) IInteger
	Between(min, max int) IInteger
}
