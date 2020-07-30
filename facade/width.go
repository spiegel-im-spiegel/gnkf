package facade

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/width"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newNormCmd returns cobra.Command instance for show sub-command
func newWidthCmd(ui *rwi.RWI) *cobra.Command {
	widthCmd := &cobra.Command{
		Use:     "width",
		Aliases: []string{"wdth", "w"},
		Short:   "Translate character width in text",
		Long:    "Translate character width in text (UTF-8 encoding only).",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			inp, err := cmd.Flags().GetString("file")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --file option"))
			}
			out, err := cmd.Flags().GetString("output")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --output option"))
			}
			form, err := cmd.Flags().GetString("width-form")
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, "Error in --width-form option"))
			}

			//Input stream
			r := ui.Reader()
			if len(inp) > 0 {
				file, err := os.Open(inp)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("file", inp)))
				}
				defer file.Close()
				r = file
			}

			//Output stream
			w := ui.Writer()
			if len(out) > 0 {
				file, err := os.Create(out)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("output", out)))
				}
				defer file.Close()
				w = file
			}

			//Run command
			if err := width.Translate(form, w, r); err != nil {
				return debugPrint(ui, errs.Wrap(err, "", errs.WithContext("file", inp), errs.WithContext("output", out)))
			}
			return nil
		},
	}
	widthCmd.Flags().StringP("file", "f", "", "path of input text file")
	widthCmd.Flags().StringP("output", "o", "", "path of output file")
	widthCmd.Flags().StringP("width-form", "w", "fold", fmt.Sprintf("width form: [%s]", strings.Join(width.FormList(), "|")))

	return widthCmd
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