package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Class data mobil (melakukan pemodelan data moobil dengan member(info detail) ID, Merk, dan Tahun)
type mobil struct {
	ID    int    `json:"id"`
	Merk  string `json:"merk"`
	Tahun int    `json:"tahun"`
}

// Variable database in memory "sementara"
var (
	database = make(map[int]mobil)
)

// JSON Respon
func setResp(res http.ResponseWriter, msg []byte, httpCode int) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(msg)
}

func main() {

	// init database
	database[1] = mobil{ID: 1, Merk: "Honda Jazz", Tahun: 2020}
	database[2] = mobil{ID: 2, Merk: "Toyota Avanza", Tahun: 2019}
	database[3] = mobil{ID: 3, Merk: "Suzuki Baleno", Tahun: 2007}

	// main route (route utama)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		msg := []byte(`"{message}" : "Server is up!"`)
		setResp(res, msg, http.StatusOK)
	})

	// route get data mobil (pemanggilan seluruh data mobil)
	http.HandleFunc("/get-mobils", func(res http.ResponseWriter, req *http.Request) {
		// Validasi method
		if req.Method != "GET" {
			msg := []byte(`"{message}" : "invalid http method"`)
			setResp(res, msg, http.StatusOK)
			return
		}
		// temp list class data mobil
		var mobils []mobil
		// loop data mobil into mobils
		for _, mobil := range database {
			mobils = append(mobils, mobil)
		}
		// encode to JSON if not error
		JSONmobils, err := json.Marshal(&mobils)
		if err != nil {
			msg := []byte(`"{message}" : "internal server error"`)
			setResp(res, msg, http.StatusInternalServerError)
			return
		}
		setResp(res, JSONmobils, http.StatusOK)
	})

	// route add mobil (menambahkan data mobil ke database)
	http.HandleFunc("/add-mobil", func(res http.ResponseWriter, req *http.Request) {
		// Validasi Method
		if req.Method != "POST" {
			msg := []byte(`"{message}" : "invalid http method"`)
			setResp(res, msg, http.StatusOK)
			return
		}
		var mobilx mobil                                // variable class mobil
		payload := req.Body                             // mengambil/req payload yang dikirim dari sisi client
		defer req.Body.Close()                          // menutup req
		err := json.NewDecoder(payload).Decode(&mobilx) // melakukan deserialize / decode dari payload ke item mobilx
		// Validasi jika data error / tidak terisi misalnya
		if err != nil {
			msg := []byte(`"{message}" : "error when parsing data"`)
			setResp(res, msg, http.StatusInternalServerError)
			return
		}
		// menambahkan data dari moobilx ke database dengan ID (sesuai format map)
		database[mobilx.ID] = mobilx
		msg := []byte(`"{message}" : "success create data mobil"`)
		setResp(res, msg, http.StatusOK)
	})

	// Config port
	err := http.ListenAndServe(":65", nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

}
