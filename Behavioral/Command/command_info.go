package Command

import (
	"fmt"
)

type Command interface {
	Info() string
}

/*
	Menghitung Waktu
*/
type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

/*
	Cetak ke Layar
*/
type HelloMessage struct{}

func (h HelloMessage) Info() string {
	return "Hello world!"
}
