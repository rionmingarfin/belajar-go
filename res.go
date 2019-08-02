package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//port
const (
	ListeningPort = ":8081"
)

//response struct
type Response struct {
	Title  string
	Detail string
}

func helloWord(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("helloword"))
	resp := Response{}
	resp.Title = "selamat datang"
	resp.Detail = "jelas online golang"
	encodedResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}

//penyimpanan
type siswa struct {
	Id    int
	Nama  string
	Kelas int
}

var SemuaSiswa []siswa
var id int

//get semua siswa
func ReadAll(w http.ResponseWriter, r *http.Request) {
	encodedResp, err := json.Marshal(SemuaSiswa)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}
func createSiswa(w http.ResponseWriter, r *http.Request) {
	// Get the Body
	siswaBaru := siswa{}
	// mengubah inputan json menjadi struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&siswaBaru)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	id++
	siswaBaru.Id = id
	SemuaSiswa = append(SemuaSiswa, siswaBaru)
	resp := Response{}
	resp.Title = "sukses"
	resp.Detail = "sudah masuk"
	encodedResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
	//caara rapinya
	// if err != nil {
	// 	resp.AddError(errorDecodingJSONReq, errorDecodingJSONReq)
	// 	sendResponse(http.StatusBadRequest, resp, w, r)
	// 	return
	// }
}

//function detail siswa
func GetDetailSiswa(w http.ResponseWriter, r *http.Request) {
	id, err := getVarsID(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	var siswatertentu siswa
	// searching siswa tertentu
	for _, perSiswa := range SemuaSiswa {
		if perSiswa.Id == id {
			siswatertentu = perSiswa
		}
	}
	//kalau siswa id tidak di temukan
	if siswatertentu.Id == 0 {
		resp := Response{}
		resp.Detail = "id tidak di temukan"
		encodedResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}

		w.Write(encodedResp)
		return
	}
	//kalau di temukan
	encodedResp, err := json.Marshal(siswatertentu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}
func updateSiswa(w http.ResponseWriter, r *http.Request) {
	// decode json body
	siswaUpdate := siswa{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&siswaUpdate)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	//temukan id
	id, err := getVarsID(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	var siswatertentu siswa
	indexDitemukan := -1
	// searching siswa tertentu
	for i, perSiswa := range SemuaSiswa {
		if perSiswa.Id == id {
			siswatertentu = perSiswa
			indexDitemukan = i
		}
	}
	//kalau siswa id tidak di temukan
	if siswatertentu.Id == 0 {
		resp := Response{}
		resp.Detail = "id tidak di temukan"
		encodedResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}

		w.Write(encodedResp)
		return
	}
	//kalau di temukan
	SemuaSiswa[indexDitemukan].Nama = siswaUpdate.Nama
	SemuaSiswa[indexDitemukan].Kelas = siswaUpdate.Kelas
	encodedResp, err := json.Marshal(siswaUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}
func deleteSiswa(w http.ResponseWriter, r *http.Request) {

	//temukan id
	id, err := getVarsID(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	var siswatertentu siswa
	indexDitemukan := -1
	// searching siswa tertentu
	for i, perSiswa := range SemuaSiswa {
		if perSiswa.Id == id {
			siswatertentu = perSiswa
			indexDitemukan = i
		}
	}
	//kalau siswa id tidak di temukan
	if siswatertentu.Id == 0 {
		resp := Response{}
		resp.Detail = "id tidak di temukan"
		encodedResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}

		w.Write(encodedResp)
		return
	}
	//kalau di temukan
	SemuaSiswa = append(SemuaSiswa[:indexDitemukan], SemuaSiswa[indexDitemukan+1:]...)
	resp := Response{}
	resp.Detail = "siswa berhasil di hapus"
	encodedResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}

	w.Write(encodedResp)
	return
}
func getVarsID(r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		convertedVal, err := strconv.Atoi(val)
		if err != nil {
			return id, err
		}
		id = convertedVal
	}
	return
}
func main() {
	//   dumy siswa
	SemuaSiswa = []siswa{}
	// siswaBaru := siswa{
	// 	Id:    1100111,
	// 	Nama:  "rion",
	// 	Kelas: 9,
	// }
	// SemuaSiswa = append(SemuaSiswa, siswaBaru)
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", helloWord).Methods(http.MethodGet)
	r.HandleFunc("/api/siswa", ReadAll).Methods(http.MethodGet)
	r.HandleFunc("/api/siswa", createSiswa).Methods(http.MethodPost)
	r.HandleFunc("/api/siswa/{id:[0-9]+}", GetDetailSiswa).Methods(http.MethodGet)
	r.HandleFunc("/api/siswa/{id:[0-9]+}", updateSiswa).Methods(http.MethodPatch)
	r.HandleFunc("/api/siswa/{id:[0-9]+}", deleteSiswa).Methods(http.MethodDelete)

	// running port
	log.Printf("Starting http server at %v", ListeningPort)
	err := http.ListenAndServe(ListeningPort, r)
	if err != nil {
		log.Fatalf("Unable to run http server: %v", err)
	}
	log.Println("Stopping API Service...")
}
