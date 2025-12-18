package utils

import "github.com/atotto/clipboard"

func WriteToClipboard(text string) error {
	return clipboard.WriteAll(text)
}
