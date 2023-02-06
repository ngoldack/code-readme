package internal

import (
	"os"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func Title() {
	// Print a large text with differently colored letters.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("code", pterm.NewStyle(pterm.FgCyan)),
		putils.LettersFromStringWithStyle("-", pterm.NewStyle(pterm.FgWhite)),
		putils.LettersFromStringWithStyle("readme", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()
	time.Sleep(time.Millisecond * 100)
	pterm.DefaultSpinner.Start()
}

func Done() {
	pterm.DefaultSpinner.Stop()
	pterm.Success.Println("Done!")
}

func Info(s string) {
	pterm.Info.Println(s)
}

func Warning(s string) {
	pterm.Warning.Println(s)
}

func Fatal(s string) {
	Error(s)
	os.Exit(1)
}

func Error(s string) {
	pterm.Error.Println(s)
}

func Input(untilValid bool, multiLine bool) string {
	input := pterm.DefaultInteractiveTextInput.WithMultiLine(multiLine)

	if untilValid {
		var result string
		for len(result) == 0 {
			result, _ = input.Show()
		}
		return result
	}

	result, _ := input.Show()
	return result
}
