package Future

import (
	"errors"
	"sync"
	"testing"
)

type MaybeString struct{}

func TestStringOrError_Execute(t *testing.T) {

	future := &MaybeString{}

	t.Run("Failed result", func(t *testing.T) {

		var wg sync.WaitGroup
		wg.Add(1)

		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})

		wg.Wait()
	})
}
