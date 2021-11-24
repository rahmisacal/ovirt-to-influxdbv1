package infrastructure

import (
	"../domain"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetToken(tokenUrl string) string {

	payload := strings.NewReader(`grant_type=password&scope=ovirt-app-api&username=admin%40internal&password=M2n!SBLz.`)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", tokenUrl, payload)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + "POST request was not generated.\n")
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic Og==")
	req.Header.Add("Cookie", "locale=en_US")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + " Response Client was not generated.\n")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		fmt.Printf("ERROR : " + time.Now().String() + "response body was not read as json.\n")
	}

	object := domain.Object{}
	_ = json.Unmarshal(body, &object)

	//fmt.Println(object.Token)
	return object.Token
}
