package main

import (
	"encoding/json"
	"fmt"
)

type statusCodes []statusCode

type statusCode struct {
	Code    int    `json:"code"`
	Descrip string `json:"descrip"`
}

func main() {
	var rcvd string

	var s statusCodes

	rcvd = `[{"code":200,"descrip":"StatusOK"},{"code":301,"descrip":"StatusOK"},{"code":302,"descrip":"StatusFound"},{"code":303,"descrip":"StatusSeeOther"},{"code":307,"descrip":"StatusTemporaryRedirect"},{"code":400,"descrip":"StatusBadRequest"},{"code":401,"descrip":"StatusUnauthorized"},{"code":402,"descrip":"StatusPaymentRequired"},{"code":403,"descrip":"StatusForbidden"},{"code":404,"descrip":"StatusNotFound"},{"code":405,"descrip":"StatusMethodNotAllowed"},{"code":418,"descrip":"StatusTeapot"},{"code":500,"descrip":"StatusInternalServerError"}]`

	error := json.Unmarshal([]byte(rcvd), &s)
	if error != nil {
		fmt.Println("error", error)
	}

	for _, status := range s {
		fmt.Printf("%+v\n", status)
	}
}
