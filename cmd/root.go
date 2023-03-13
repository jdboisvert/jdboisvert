/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"

	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func handleSelection(result string) {
	switch result {
	case "Send Email":
		fmt.Println("Opening your default email client...")
		cmd := exec.Command("open", "mailto:info.jeffreyboisvert@gmail.com")

		// Run the command and check for errors
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}

	case "Play a Game":
		fmt.Println("Alright lets play...")

	case "Quit":
		fmt.Println("Bye and have a great day!")
	}
}

var rootCmd = &cobra.Command{
	Use:   "jdboisvert",
	Short: "A personal card for Jeffrey Boisvert (@jdboisvert)",
	Long:  "jdboisvert is a personal card for Jeffrey Boisvert (@jdboisvert). This is a CLI application helps you to get to know me a little better and be able to better reach me.",
	Run: func(cmd *cobra.Command, args []string) {
		Box := box.New(box.Config{Px: 4, Py: 1, Type: "Double", Color: "Cyan", ContentAlign: "Center"})
		Box.Print(
			"Jeffrey Boisvert (@jdboisvert)",
			"Useful links:\n"+
				"Linkedin: https://www.linkedin.com/in/jeffreybv/\n"+
				"DEV: https://dev.to/jdvert\n"+
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

		quitText := "Quit"
		sendEmailText := "Send Email"
		downloadResumeText := "Play a Game"

		prompt := promptui.Select{
			Label: "So what do you want to do now?",
			Items: []string{sendEmailText, downloadResumeText, quitText},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		handleSelection(result)

		for result != "Quit" {
			prompt = promptui.Select{
				Label: "So what do you want to do now?",
				Items: []string{sendEmailText, downloadResumeText, quitText},
			}

			_, result, err = prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			handleSelection(result)

		}
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
