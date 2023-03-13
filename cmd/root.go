/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jdboisvert",
	Short: "A personal card for Jeffrey Boisvert (@jdboisvert)",
	Long:  "jdboisvert is a personal card for Jeffrey Boisvert (@jdboisvert). This is a CLI application helps you to get to know me a little better.",
	Run: func(cmd *cobra.Command, args []string) {
		Box := box.New(box.Config{Px: 2, Py: 5, Type: "Double", Color: "Green", ContentAlign: "Center"})
		Box.Print(
			"Jeffrey Boisvert (@jdboisvert)",
			"Current Job: Staff Software Developer at Newton Crypto\n"+
				"\n\n"+
				"The Socials:\n"+
				"Linkedin: https://www.linkedin.com/in/jeffreybv/\n"+
				"DEV: https://dev.to/jdvert\n"+
				"\n\n"+
				"The Code/Projects:\n"+
				"PyPi: https://pypi.org/user/jdboisvert/\n"+
				"GitHub: https://github.com/jdboisvert\n"+
				"\n\n"+
				"As the famous Alan Turing once said: \n"+
				"\"We can only see a short distance ahead, but\n"+
				" we can see plenty there that needs to be done.\""+
				"\n\n"+
				"I believe there is always something to improve on and build upon.\n"+
				"I am always looking for new challenges \n"+
				"and opportunities to learn and grow."+
				"\n\n"+
				"Thanks for checking out my card! Happy coding!",
		)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jdboisvert.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
