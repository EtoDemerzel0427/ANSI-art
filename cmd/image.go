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
	"ANSI-art/ansi"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"log"
)

var (
	ImgWidth int
	ImgHeight int
	Sequence string
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Show your image in the terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		src, err := imaging.Open(filename)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		src = imaging.Resize(src, ImgWidth, ImgHeight, imaging.Lanczos)
		fmt.Print(ansi.Pixels2ColoredANSI(src, Sequence))
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)

	imageCmd.Flags().IntVarP(&ImgWidth, "width", "W", 100, "the resized width of the image")
	imageCmd.Flags().IntVarP(&ImgHeight, "height", "H", 100, "the resized height of the image")
	imageCmd.Flags().StringVarP(&Sequence, "seq", "s",
		"01", "the string of ANSI chars that build the image")



}
