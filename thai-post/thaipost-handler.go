package thai_post

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

var Token string

func init() {
	tokenResp, err := getToken()

	if err != nil {
		log.Error(err)
	}

	Token = tokenResp.Token
	println(Token)
}

func GetEmsInformation(message string) string {
	barcode := []string{}
	barcode = append(barcode, message)

	item, err := getItemsDetail(barcode)
	if err != nil {
		log.Error(err)
	}
	return fmt.Sprintf("%v", item.Response.Items)
}
