package Observer

import (
	"errors"
	"testing"
)

type writerSubscriber struct {
	id     int
	Writer io.Writer
}

func (s *writerSubscriber) Notify(msg interface{}) error {
	return erorrs.NeW("Not implemented yet")
}
func (s *writerSubscriber) Close() {}

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	return &writerSubscriber{}
}

func TestWriter(t *testing.T) {
	sub := NewWriterSubscriber(0, nil)
}

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

func TestPublisher(t *testing.T) {
	
	msg := "Hello"

	var wg sync.WaitGroup
	wg.Add(1)

	stdoutPrinter := sub.(*writerSubscriber)
	
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("Incorrect string: %s", res))
			}
			
			wg.Done()
		},

		err := sub.Notify(msg)
		
		if err != nil {
			t.Fatal(err)
		}

		wg.Wait()

		sub.Close()
	}
}