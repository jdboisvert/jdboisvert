/*
Copyright © 2023 Jeffrey Boisvert info.jeffreyboisvert@gmail.com
*/
package cmd

import (
	"os"
	"os/exec"

	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/jdboisvert/jdboisvert/pkg"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	SendEmailText  = "Send Email"
	TakeAQuizText  = "Take a Quiz"
	QuitText       = "Quit"
	PromptQuestion = "So what do you want to do now?"
	ErrorText      = "So it seems you may have run into an issue :(. Please reach out and I can see how I can help you out!"
)

func handleSelection(result string) {
	switch result {
	case SendEmailText:
		fmt.Println("Opening your default email client...")
		cmd := exec.Command("open", "mailto:info.jeffreyboisvert@gmail.com")

		// Run the command and check for errors
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}

	case TakeAQuizText:
		fmt.Println("Alright good luck...")
		pkg.QuizGameStart()

	case QuitText:
		fmt.Println("Bye and have a great day!")
	}
}

func getPromptResult() (string, error) {
	prompt := promptui.Select{
		Label: PromptQuestion,
		Items: []string{SendEmailText, TakeAQuizText, QuitText},
	}
	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func promptForSelection() {
	// Prompt the user at least once
	result, err := getPromptResult()

	if err != nil {
		fmt.Println(ErrorText)
		return
	}

	handleSelection(result)

	for result != QuitText {
		result, err = getPromptResult()

		if err != nil {
			fmt.Println(ErrorText)
			return
		}

		handleSelection(result)
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

		promptForSelection()
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
