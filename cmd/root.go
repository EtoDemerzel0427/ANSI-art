package cmd

import (
	"github.com/spf13/cobra"
)


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ansi-art",
	Short: "An ANSI-art style img & Gif viewer.",
	Long: `ANSI-Art converts your image & Gif to ANSI characters (with color) that can show in the console. 
Users can speicify the characters to use, size of the image, etc to get the best results. some sample commands are:

./ansi-art image pic/messi.png -W 150 -H 60 -a -C 40
./ansi-art gif pic/sharingan.gif -W 150 -H 60 -s SASUKE -d 300 -m bgm/uefa.mp3`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

