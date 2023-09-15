package helpers

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func ClientRequest(method string, url string, reqBody interface{}, headers interface{}) []byte {
	client := httpClient()
	endpoint := url
	values := reqBody
	jsonData, err := json.Marshal(values)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	headersMap := headers.(map[string]string)
	for key, val := range headersMap {
		req.Header.Add(key, val)
	}

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body
}
