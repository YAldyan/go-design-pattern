package Barrier

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Err  error
	Resp string
}

func barrier(endpoints ...string) {

	// jumlah request sesuai dengan jumlah endpoint
	requestNumber := len(endpoints)

	// buat channel tipe barrierResp dan Jumlah Buffer maksimal sesuai jumlah endpoint
	in := make(chan barrierResp, requestNumber)

	// akan dieksekusi ketika semua code telah dijalankan
	defer close(in)

	/*
		membuat variabel bukan channel dengan
		tipe BarrierResp dan panjang array
		adalah sesuai dengan requestNumber
	*/
	responses := make([]barrierResp, requestNumber)
	for _, endpoint := range endpoints {

		/*
			jalankan request ke endpoint

			channel in akan diisi nilainya
			ketika fungsi makeRequest jalan
			sesuai dengan response yang diperoleh
		*/
		go makeRequest(in, endpoint)
	}

	var hasError bool

	/*
		lakukan iterasi dan simpan respon yang didapat
		apakah itu respon valid atau malah error
	*/
	for i := 0; i < requestNumber; i++ {

		/*
			akan dilakukan pengiriman data dari Channel in
			karena menggunakan Buffered Channel maka, hanya
			ada 2, jadi pengiriman akan dilakukan iterasi
			dengan menggunakan for, karena tiap kali data
			dikirim, maka akan berkurang, hingga tidak ada
			data yang tersisa d channel
		*/
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}

		/*
			simpan tiap-tiap respon yang diperoleh
			sesuai dengan jumlah request yang diek-
			sekusi sebelumnya.
		*/
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

// out channel tipenya adanya channel penerima
func makeRequest(out chan<- barrierResp, url string) {

	// buat instance barrierResp
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) *
			time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}
