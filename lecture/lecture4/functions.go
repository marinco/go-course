package lecture4

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
	}
	fmt.Println("Status:", resp.Status)
	fmt.Println("Status code:", resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println("Body:", string(body))
}
