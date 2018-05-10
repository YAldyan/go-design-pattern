package Shapes

import (
	"Strategy"
	"fmt"
)

type TextSquare struct {
	Strategy.PrintOutput
}

func (t *TextSquare) Print() error {
	r := bytes.NewReader([]byte("Circle"))
	io.Copy(t.Writer, r)
	return nil
}
