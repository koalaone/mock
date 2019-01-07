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

type IString interface {
	Raw() string
	Path(path string) IValue
	Schema(value interface{}) IString
	Equal(value interface{}) IString
	NotEqual(value interface{}) IString
	Null() IString
	NotNull() IString
	EqualFold(value string) IString
	NotEqualFold(value string) IString
	Length() IInteger
	Datetime(layout string) IDatetime
	EqualEmpty() IString
	NotEqualEmpty() IString
	Contains(value interface{}) IString
	NotContains(value interface{}) IString
	ContainFold(value interface{}) IString
	NotContainFold(value interface{}) IString
	Match(expr string) IMatch
	MatchMulti(expr string) []IMatch
	NotMatch(expr string) IString
}
