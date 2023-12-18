package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"crypto/tls"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    response, err := http.Get("https://www.cloudflarestatus.com/api/v2/components.json")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))

}