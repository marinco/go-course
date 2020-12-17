package lecture4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "lecture/lecture4/index.html")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/file" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	var path = "lecture/lecture4/" + r.URL.Query().Get("url")
	w.Header().Set("Content-Disposition", "attachment; filename="+path)
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, path)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("POST response!\n"))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/json" && r.Method != http.MethodPost {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Println("body is ", string(body))

	var person Person
	json.Unmarshal(body, &person)

	person.Status = "PROCESSED"

	response, _ := json.Marshal(person)

	w.Header().Set("Content-type", "application/json")
	w.Write(response)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "custom error: %d \n", status)
}
