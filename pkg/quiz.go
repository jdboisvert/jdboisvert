package pkg

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type Question struct {
	Text       string
	Options    []string
	CorrectIdx int
}

func GenerateQuestions() []Question {
	return []Question{
		{
			Text:       "What is my first name?",
			Options:    []string{"Jeffrey", "Jeff", "Alan", "Python"},
			CorrectIdx: 0,
		},
		{
			Text:       "Where was I working January 23rd, 2023?",
			Options:    []string{"Enzyme Testing Labs", "Greencopper", "Newton Crypto", "Dawson College"},
			CorrectIdx: 2,
		},
		{
			Text:       "What language is this program written in (hint checkout https://github.com/jdboisvert/jdboisvert)?",
			Options:    []string{"Rust", "Python", "Javascript", "Go"},
			CorrectIdx: 3,
		},
		{
			Text:       "I have a module publish on PyPi called django-migration-<your response>",
			Options:    []string{"previous", "rollback", "migrate", "zero"},
			CorrectIdx: 1,
		},
	}
}

func HandleResult(score int, total int) {
	fmt.Printf("You got %d out of %d questions correct!\n", score, total)

	decimalScore := float64(score) / float64(total)
	if decimalScore < 0.5 {
		fmt.Println("Better luck next time! Why not send me an email to get to know me better!")
	} else if decimalScore < 0.75 {
		fmt.Println("Not bad! Why not send me an email to get to know me better!")
	} else {
		fmt.Println("Perfect! Send me an email so I can now get to know you better!")
	}
}

func PlayQuiz(questions []Question) int {
	numberOfQuestionsCorrect := 0
	for _, question := range questions {
		prompt := promptui.Select{
			Label: question.Text,
			Items: question.Options,
		}
		index, _, err := prompt.Run()
		if err != nil {
			// Error exit
			return 0
		}

		if index == question.CorrectIdx {
			numberOfQuestionsCorrect++
		}
	}
	return numberOfQuestionsCorrect
}

func QuizGameStart() {
	questions := GenerateQuestions()
	numberOfQuestionsCorrect := PlayQuiz(questions)

	HandleResult(numberOfQuestionsCorrect, len(questions))
}
