package Barrier

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

/*
	We have a single test that will execute three subtests:

	1. The first test makes two calls to the correct endpoints
	2. The second test will have an incorrect endpoint, so it must return an error
	3. The last test will return the maximum timeout time so that we can force a timeout error
*/

func TestBarrier(t *testing.T) {

	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Accept-Encoding") || strings.Contains(result, "User-Agent") {
			t.Fail()
		}

		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "ERROR") {
			t.Fail()
		}

		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		// var timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}

		t.Log(result)
	})
}

// func barrier(endpoints ...string) {}

func captureBarrierOutput(endpoints ...string) string {

	/*
		pipe allows us to connect an io.Writer interface to an io.Reader interface
		so that the reader input is the Writer output
	*/
	reader, writer, _ := os.Pipe()

	/*
		We define the os.Stdout as the writer
	*/
	os.Stdout = writer
	out := make(chan string)

	/*
		Then, to capture stdout output, we will need a different Goroutine that listens while we write to the console.

		As you know, if we write, we don't capture, and if we capture, we are not writing. The keyword here is while;

		it is a good rule of thumb that if you find this word in some definition, you'll probably need a concurrent structure.

		So we use the go keyword to launch a different Goroutine that copies reader input to a bytes buffer before sending the
		contents of the buffer through a channel (that we should have previously created).

		Go Routine ini bersifat blocking karena unbuffered, sehingga Main Program akan menunggu serah terima objek baru bisa
		lanjutkan ke proses berikutnya
	*/
	go func() {
		var buf bytes.Buffer

		/*
			nilai reader akan berubah sesuai dengan nilai writernya
			dikarenakan writer dan reader sudah disinkronisasikan.
			ketika writer berisi nilai dari eksekusi fungsi barrier
			maka nilainya akan dicopy ke buffer dan nantinya di kirim
			ke main program untuk dicetak ke layar.
		*/
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	writer.Close()
	temp := <-out
	return temp
}
