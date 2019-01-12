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

package reporter

import "testing"

func TestTestingReporterErrorFormat(t *testing.T) {
	tr := NewTestingReporter(t)
	tr.ErrorFormat("report error value:%v", "test")

	tr.ErrorFormat("report error value:%v", "test123")
}

func TestTestingReporterInfoFormat(t *testing.T) {
	tr := NewTestingReporter(t)
	tr.InfoFormat("report info value:%v", "test")

	tr.InfoFormat("report info value:%v", "test123")
}
