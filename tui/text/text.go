package text

import "fmt"

type Text struct {
	text string
}

func New(text string) *Text {
	return &Text{text}
}

func (t Text) String() string {
	return t.text
}

func (t *Text) Bold() *Text {
	t.text = fmt.Sprintf("\x1b[1;39m%s\x1b[0m", t.text)
	return t
}
