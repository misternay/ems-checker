package thai_post

import (
	"bytes"
	"crypto/tls"
	"ems-checker/thai-post/model"
	"encoding/json"
	"errors"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func getToken() (model.GetToken, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Timeout:   time.Duration(5 * time.Second),
		Transport: tr,
	}

	request, err := http.NewRequest(http.MethodPost, "https://trackapi.thailandpost.co.th/post/api/v1/authenticate/token", nil)

	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", os.Getenv("AUTH"))

	if err != nil {
		log.Error(err)
	}

	resp, error := client.Do(request)

	if error != nil {
		log.Error(error)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
	}

	tokenResp := model.GetToken{}

	if strings.TrimSpace(resp.Status) != "200" {
		return tokenResp, errors.New("TrackAndTrace thailandpost have a problem now" + resp.Status)
	}

	err = json.Unmarshal(body, &tokenResp)

	if err != nil {
		log.Error(err)
	}

	return tokenResp, nil
}

func getItemsDetail(barcode []string) (model.GetItems, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Timeout:   time.Duration(30 * time.Second),
		Transport: tr,
	}

	reqBody := &model.ServiceRequest{
		Barcode:  barcode,
		Language: "TH",
		Status:   "all",
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(reqBody)

	if err != nil {
		log.Error(err)
	}

	request, err := http.NewRequest(http.MethodPost, "https://trackapi.thailandpost.co.th/post/api/v1/track", buf)

	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", "Token "+Token)

	if err != nil {
		log.Error(err)
	}

	resp, error := client.Do(request)

	if error != nil {
		log.Error(error)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
	}

	items := model.GetItems{}

	if strings.TrimSpace(resp.Status) != "200" {
		return items, errors.New("TrackAndTrace thailandpost have a problem now" + resp.Status)
	}

	err = json.Unmarshal(body, &items)

	if err != nil {
		log.Error(err)
	}
	println(items.Status)
	return items, nil
}
