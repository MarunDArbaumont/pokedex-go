package repl

import "strings"

func CleanInput(text string) []string {
        text = strings.ToLower(text)
        splitedText := strings.Split(text, " ")
        return splitedText
}
