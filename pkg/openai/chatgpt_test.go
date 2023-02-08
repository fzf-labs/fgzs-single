package openai

import (
	"fmt"
	"testing"
)

func TestChatGPT_Completions(t *testing.T) {
	newChatGPT := NewChatGPT("")
	completions, err := newChatGPT.Completions("互联网中常见的运营指标有哪些?")
	if err != nil {
		return
	}
	fmt.Println(completions)
}
