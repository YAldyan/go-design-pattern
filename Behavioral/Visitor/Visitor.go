package visitor

import (
	"fmt"
	"io"
	"os"
)

/*Objek Yang akan dikunjungi*/
type MessageA struct {
	Msg    string
	Output io.Writer
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}

/*Objek yang bisa kunjungi*/
type Visitable interface {
	Accept(Visitor)
}

func (m *MessageA) Accept(v Visitor) {
	v.VisitA(m)
}
func (m *MessageB) Accept(v Visitor) {
	v.VisitB(m)
}

/* Kelas/Objek yang akan mengunjungi*/
type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type MessageVisitor struct{}

func (mf *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}
func (mf *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

type MsgFieldVisitorPrinter struct{}

func (mf *MsgFieldVisitorPrinter) VisitA(m *MessageA) {
	fmt.Printf(m.Msg)
}
func (mf *MsgFieldVisitorPrinter) VisitB(m *MessageB) {
	fmt.Printf(m.Msg)
}
