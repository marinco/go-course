package lecture4

import (
	"net/http"
)

func Main() {
	//get("https://google.com")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/redirect", redirectHandler)

	http.ListenAndServe(":8080", nil)
}
