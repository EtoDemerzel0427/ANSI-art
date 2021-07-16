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
	"time"
)

//var GifFilename string
var (
	gifWidth int
	gifHeight int
	duration int
	gifFile string
	gifSeq string
)

// gifCmd represents the gif command
var gifCmd = &cobra.Command{
	Use:   "gif",
	Short: "Playing gif in your terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		decode.Gif2imgs(gifFile, gifWidth, gifHeight, time.Duration(duration*1000000), "ROSE")
	},
}

func init() {
	rootCmd.AddCommand(gifCmd)

	gifCmd.Flags().StringVarP(&gifFile, "filename", "f", "demo.gif", "the input gif file")
	gifCmd.Flags().IntVarP(&gifWidth, "width", "W", 100, "the resized width of the image")
	gifCmd.Flags().IntVarP(&gifHeight, "height", "H", 100, "the resized height of the image")
	gifCmd.Flags().IntVarP(&duration, "duration", "d", 200, "the duration(ms) of each frame, used to control speed")
	gifCmd.Flags().StringVarP(&gifSeq, "seq", "s", "01", "the string of ANSI chars that build the image")
}
