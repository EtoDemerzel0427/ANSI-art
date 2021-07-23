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
	"ANSI-art/ascii"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"log"
)

var (
	imgWidth     int
	imgHeight    int
	imgSeq       string
	imgFile      string
	blockMode    bool
	imgContrast  float64
	imgAsciiMode bool
	imgSigma     float64
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Show your image in the terminal.",
	Run: func(cmd *cobra.Command, args []string) {

		src, err := imaging.Open(imgFile)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}

		if imgContrast < -100. {
			imgContrast = -100.
		}
		if imgContrast > 100. {
			imgContrast = 100.
		}
		src = imaging.Resize(src, imgWidth, imgHeight, imaging.Lanczos)
		src = imaging.AdjustContrast(src, imgContrast)
		src = imaging.Sharpen(src, imgSigma)
		fmt.Print(ansi.ClearScreen())

		if imgAsciiMode {
			src = imaging.Grayscale(src)
			fmt.Println(ascii.Pixels2Ascii(src))

		} else {
			if blockMode {
				fmt.Println(ansi.Pixels2ColoredBlocks(src))
			} else {
				fmt.Println(ansi.Pixels2ColoredANSI(src, imgSeq))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.Flags().StringVarP(&imgFile, "filename", "f", "demo.gif", "the input gif file")
	imageCmd.Flags().BoolVarP(&imgAsciiMode, "ascii", "a", false, "ansi or ascii art")
	imageCmd.Flags().BoolVarP(&blockMode, "blockMode", "b", false, "character or block mode")
	imageCmd.Flags().IntVarP(&imgWidth, "width", "W", 100, "the resized width of the image")
	imageCmd.Flags().IntVarP(&imgHeight, "height", "H", 100, "the resized height of the image")
	imageCmd.Flags().Float64VarP(&imgContrast, "contrast", "C", 0., "increase/decrease the imgContrast (-100 ~ 100)")
	imageCmd.Flags().Float64VarP(&imgSigma, "sigma", "S", 0., "sharpening factor")
	imageCmd.Flags().StringVarP(&imgSeq, "seq", "s",
		"01", "the string of ANSI chars that build the image")



}
