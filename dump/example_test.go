package dump_test

import (
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/gnkf/dump"
)

func ExampleOctet() {
	dump.Octet(os.Stdout, strings.NewReader("こんにちは，世界！\n私の名前は Spiegel です。"))
	//Output:
	//0xe3, 0x81, 0x93, 0xe3, 0x82, 0x93, 0xe3, 0x81, 0xab, 0xe3, 0x81, 0xa1, 0xe3, 0x81, 0xaf, 0xef, 0xbc, 0x8c, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c, 0xef, 0xbc, 0x81, 0x0a, 0xe7, 0xa7, 0x81, 0xe3, 0x81, 0xae, 0xe5, 0x90, 0x8d, 0xe5, 0x89, 0x8d, 0xe3, 0x81, 0xaf, 0x20, 0x53, 0x70, 0x69, 0x65, 0x67, 0x65, 0x6c, 0x20, 0xe3, 0x81, 0xa7, 0xe3, 0x81, 0x99, 0xe3, 0x80, 0x82
}

func ExampleUnicodePoint() {
	dump.UnicodePoint(os.Stdout, strings.NewReader("こんにちは，世界！\n私の名前は Spiegel です。"))
	//Output:
	//0x3053, 0x3093, 0x306b, 0x3061, 0x306f, 0xff0c, 0x4e16, 0x754c, 0xff01, 0x000a, 0x79c1, 0x306e, 0x540d, 0x524d, 0x306f, 0x0020, 0x0053, 0x0070, 0x0069, 0x0065, 0x0067, 0x0065, 0x006c, 0x0020, 0x3067, 0x3059, 0x3002
}

/* Copyright 2020 Spiegel
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