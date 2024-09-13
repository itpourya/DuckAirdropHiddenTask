package process

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

const (
	contentType = "application/json"
	API_URL     = "https://api.apiduck.xyz/user-partner-mission/claim"
)

func DoProcess(ID int, wg *sync.WaitGroup, authurization string) {
	defer wg.Done()

	payload, err := json.Marshal(
		map[string]int{
			"partner_mission_id": ID,
		},
	)
	if err != nil {
		log.Fatalln("MARSHAL PAYLOAD ERROR")
	}

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("API CONNECTION ERROR")
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", authurization)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("CANT PROCESS REQUEST")
	}

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, res.Body)
	if err != nil {
		log.Fatal("getTaps: ", err.Error())
	}

	str := buffer.String()

	fmt.Println(str)
}
