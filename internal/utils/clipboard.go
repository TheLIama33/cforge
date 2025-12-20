package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/atotto/clipboard"
)

func WriteToClipboard(text string) error {
	err := clipboard.WriteAll(text)
	if err == nil {
		return nil
	}

	b64 := base64.StdEncoding.EncodeToString([]byte(text))
	osc52Sequence := fmt.Sprintf("\x1b]52;c;%s\x07", b64)

	fmt.Print(osc52Sequence)

	return nil
}
