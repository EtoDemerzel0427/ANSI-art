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
	"github.com/EtoDemerzel0427/ANSI-art/art"
	"github.com/EtoDemerzel0427/ANSI-art/decode"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

//var GifFilename string
var (
	musicFile string
	gifWidth int
	gifHeight int
	duration int
	gifFile string
	gifSeq string
	loopNum int
	gifMode bool
	gifContrast float64
	gifAsciiMode bool
	gifSigma float64
)

// gifCmd represents the gif command
var gifCmd = &cobra.Command{
	Use:   "gif",
	Short: "Playing gif in your terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(musicFile)
		if err != nil {
			log.Fatal(err)
		}
		if gifContrast < -100. {
			gifContrast = -100.
		}
		if gifContrast > 100. {
			gifContrast = 100.
		}

		var mode = art.AsciiText
		if !gifAsciiMode {
			if blockMode {
				mode = art.AnsiBlock
			} else {
				mode = art.AnsiText
			}
		}

		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		defer streamer.Close()

		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

		ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
		speaker.Play(ctrl)

		done := make(chan bool)
		as := art.NewSolver(gifWidth, gifHeight, gifContrast, gifSigma, gifSeq, mode)
		gd := decode.NewGifDecoder(as, loopNum, time.Duration(duration*1000000), &done)
		go gd.Decode(gifFile)
		//go decode.Gif2imgs(gifFile, gifWidth, gifHeight, time.Duration(duration*1000000), gifSeq, loopNum,
		//	gifAsciiMode, gifContrast, gifSigma, gifMode, &done)

		if <-done {
			speaker.Lock()
			ctrl.Paused = !ctrl.Paused
			speaker.Unlock()
		}
	},
}

func init() {
	rootCmd.AddCommand(gifCmd)
	gifCmd.Flags().BoolVarP(&gifMode, "blockMode", "b", false, "character or block mode")
	gifCmd.Flags().BoolVarP(&gifAsciiMode, "art", "a", false, "ansi or art art")
	gifCmd.Flags().StringVarP(&gifFile, "filename", "f", "pic/demo.gif", "the input gif file")
	gifCmd.Flags().StringVarP(&musicFile, "music", "m", "bgm/smb.mp3", "the background music file")
	gifCmd.Flags().IntVarP(&gifWidth, "width", "W", 100, "the resized width of the image")
	gifCmd.Flags().IntVarP(&gifHeight, "height", "H", 100, "the resized height of the image")
	gifCmd.Flags().IntVarP(&loopNum, "loop", "L", 1, "The loop number of the gif")
	gifCmd.Flags().IntVarP(&duration, "duration", "d", 200, "the duration(ms) of each frame, used to control speed")
	gifCmd.Flags().Float64VarP(&gifContrast, "contrast", "C", 0., "increase/decrease the imgContrast (-100 ~ 100)")
	gifCmd.Flags().Float64VarP(&gifSigma, "sigma", "S", 0., "sharpening factor")
	gifCmd.Flags().StringVarP(&gifSeq, "seq", "s", "01", "the string of ANSI chars that build the image")
}
