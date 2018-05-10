package Template

import (
	"strings"
)

type MessageRetriever interface {
	Message() string
}

/*
	Pola Operasinya sebagaimana
	abstrak kelas dibawah.
*/
type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

type TemplateImpl struct{}

func (t *TemplateImpl) first() string {
	return "hello"
}

func (t *TemplateImpl) third() string {
	return "template"
}

type TestStruct struct {
	Template
}

func (m *TestStruct) Message() string {
	return "world"
}
