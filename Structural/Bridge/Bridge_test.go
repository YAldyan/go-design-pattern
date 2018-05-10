package Bridge

import (
	"testing"
)

/*
	Fungsi Abstract PrinterAPI terus berubah
*/

/*
	PrinterAPI implementation standar, tanpa modifikasi
*/
func TestPrintAPI1(t *testing.T) {
	api1 := PrinterImpl1{}
	err := api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n",
			err.Error())
	}
}

/*
	Implementasi PrinterAPI telah dimodifikasi dengan kirim
	lewat io.writer, yang awalnya hanya print ke console
*/
func TestPrintAPI2(t *testing.T){

	api2 := PrinterImpl2{}
	err := api2.PrintMessage("Hello")

	if err != nil {
		expectedErrorMessage := "You need to pass an io.Writer to PrinterImpl2"

		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Error message was not correct.\n
				Actual: %s\nExpected: %s\n", err.Error(), expectedErrorMessage)
		}
	}

	testWriter := TestWriter{}
	api2 = PrinterImpl2{
		Writer: &testWriter,
	}

	expectedMessage := "Hello"
	err = api2.PrintMessage(expectedMessage)
	
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer. \n Actual:
				%s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {

	/*
		Testing untuk printer yang menampilkan pesan
		ke layar.

		Implementasi PrintAPI1
	*/
	expectedMessage := "Hello io.Writer"
	
	normal := NormalPrinter{
		Msg:expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	/*
		Implementasi Printer dengan metode
		menggunakan TestWriter.	

		Implementasi PrintAPI2
	*/
	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
			
		Printer: &PrinterImpl2{
			Writer:&testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n
			Actual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: Hello io.Writer"

	/*
		Testing untuk printer yang menampilkan pesan
		ke layar.

		Implementasi PrintAPI1
	*/
	packt := PacktPrinter{
		Msg:passedMessage,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}		

	/*
		Implementasi Printer dengan metode
		menggunakan TestWriter.	

		Implementasi PrintAPI2
	*/
	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: passedMessage,
		Printer:&PrinterImpl2{
			Writer:&testWriter,
		},
	}

	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\n
			Actual: %s\nExpected: %s\n", testWriter.Msg,expectedMessage)
	}
}