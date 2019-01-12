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

import "testing"

func Test_convertString(t *testing.T) {
	if convertString("") != "" {
		t.Errorf("convertString error ")
		return
	}

	if convertString("aa") != "aa" {
		t.Errorf("convertString error ")
		return
	}

	if convertString(1234) != "1234" {
		t.Errorf("convertString error :%v", convertString(1234))
		return
	}

	if convertString(1234.14) != "1234.14" {
		t.Errorf("convertString error :%v", convertString(1234.14))
		return
	}

	if convertString(nil) != "<nil>" {
		t.Errorf("convertString error :%v", convertString(nil))
		return
	}
}
