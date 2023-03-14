package pkg

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// captureOutput captures the output of the given function by temporarily
// redirecting stdout to a buffer.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

func TestGenerateQuestions(t *testing.T) {
	questions := GenerateQuestions()

	assert.Equal(t, len(questions), 4)

	// Check that each question has at least 2 options
	for _, q := range questions {
		assert.GreaterOrEqual(t, len(q.Options), 2)
	}

	// Check that each question has a valid index for the correct answer
	for _, q := range questions {
		assert.GreaterOrEqual(t, q.CorrectIdx, 0)
		assert.Less(t, q.CorrectIdx, len(q.Options))
	}
}

func TestHandleResult(t *testing.T) {
	t.Run("Low Score", func(t *testing.T) {
		output := captureOutput(func() {
			HandleResult(1, 4)
		})

		assert.Contains(t, output, "You got 1 out of 4 questions correct!")
		assert.Contains(t, output, "Better luck next time! Why not send me an email to get to know me better!")
	})

	t.Run("Medium Score", func(t *testing.T) {
		output := captureOutput(func() {
			HandleResult(2, 4)
		})

		assert.Contains(t, output, "You got 2 out of 4 questions correct!")
		assert.Contains(t, output, "Not bad! Why not send me an email to get to know me better!")
	})

	t.Run("High Score", func(t *testing.T) {
		output := captureOutput(func() {
			HandleResult(4, 4)
		})

		assert.Contains(t, output, "You got 4 out of 4 questions correct!")
		assert.Contains(t, output, "Perfect! Send me an email so I can now get to know you better!")
	})
}
