package Command

import (
	"fmt"
)

type Command interface {
	Execute()
}

/*
	Cetak Ke Layar
*/
type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

/*
	Create Command

	Return Comman Interface implement
	ConsoleOutput
*/
func CreateCommand(s string) Command {
	fmt.Println("Creating command")
	return &ConsoleOutput{
		message: s,
	}
}

/*
	Tipe Queue untuk menyimpan
	semua tipe yang implement
	Command Interface
*/
type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)
	if len(p.queue) == 3 {
		for _, command := range p.queue {

			/*
				Jalankan Perintah yang ada di
				dalam queue, menggunakan inter
				face Command, tapi yang akan
				dijalankan adalah Objek yang
				mengimplement Interface Command
				sesuai dengan isi dari Queue
			*/
			command.Execute()
		}
		p.queue = make([]Command, 3)
	}
}
