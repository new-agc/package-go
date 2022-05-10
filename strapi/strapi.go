package strapi

import (
	"bytes"
	"net/http"
	"encoding/json"
)

func Strapi_PostRequest (argURL string, argReqData []byte) {

	url := argURL
	postRest := argReqData

	responseBody := bytes.NewBuffer(postRest)
	resp, error := http.Post(url, "application/json", responseBody)
	//Handle Error
	if error != nil {
		log.Fatalf("An Error Occured %v", error)
	}
	defer resp.Body.Close()
}