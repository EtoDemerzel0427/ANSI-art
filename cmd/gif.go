/*
Copyright © 2021 Weiran Huang <huangweiran1998@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"ANSI-art/decode"
	"github.com/spf13/cobra"
)

//var GifFilename string

// gifCmd represents the gif command
var gifCmd = &cobra.Command{
	Use:   "gif",
	Short: "Playing gif in your terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		decode.Gif2imgs(filename)
	},
}

func init() {
	rootCmd.AddCommand(gifCmd)

	//gifCmd.Flags().StringVarP(&GifFilename, "filename", "f",
	//	"demo.gif", "The input gif file.")
}