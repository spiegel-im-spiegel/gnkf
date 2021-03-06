package kana_test

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gnkf/kana"
)

func ExampleConvertString() {
	txt := "あいうえおわゐゑをんゔゕゖゝゞアイウエオワヰヱヲンヴヵヶヽヾ"
	str, err := kana.ConvertString(kana.Hiragana, txt, false)
	if err != nil {
		return
	}
	fmt.Println(str)
	str, err = kana.ConvertString(kana.Katakana, txt, false)
	if err != nil {
		return
	}
	fmt.Println(str)
	str, err = kana.ConvertString(kana.Chokuon, txt, false)
	if err != nil {
		return
	}
	fmt.Println(str)
	//Output:
	//あいうえおわゐゑをんゔゕゖゝゞあいうえおわゐゑをんゔゕゖゝゞ
	//アイウエオワヰヱヲンヴヵヶヽヾアイウエオワヰヱヲンヴヵヶヽヾ
	//あいうえおわゐゑをんゔかけゝゞアイウエオワヰヱヲンヴカケヽヾ
}

/* Copyright 2020-2021 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
