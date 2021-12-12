//nolint
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func req() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://xx.x.com/api", nil)
	request.Header.Set("X-Header1", "Value1")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		fmt.Println("==200")
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println("==non 200")
	}
}
