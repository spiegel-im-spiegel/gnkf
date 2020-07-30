package enc

import (
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

//Decode translates from UTF-8 encodeing text.
func Decode(writer io.Writer, ianaName string, txt io.Reader) error {
	decoder, err := GetEncoding(ianaName)
	if err != nil {
		return errs.WrapWithCause(err, nil, errs.WithContext("ianaName", ianaName))
	}
	return decode(decoder, writer, txt)
}

func decode(decoder encoding.Encoding, writer io.Writer, txt io.Reader) error {
	if decoder == unicode.UTF8 {
		return notTranslate(writer, txt)
	}
	if _, err := io.Copy(writer, decoder.NewDecoder().Reader(txt)); err != nil {
		return errs.WrapWithCause(ecode.ErrInvalidEncoding, err)
	}
	return nil
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