package infrastructure

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GetMetric(url string, token string) []byte {

	var bearer = "Bearer " + token
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + "GET request was not generated.\n")
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Cookie", "locale=en_US")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + " Response Client was not generated.\n")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + "response body was not read as json.\n")
	}
	//fmt.Println(string(body))
	return body
}
