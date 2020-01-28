package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type code struct {
	Code    int    `json: "Code"`
	Descrip string `json: Descrip`
}

type codes []code

func main() {
	var data codes
	rcvd := `[{"Code": 200,"Descrip": "StatusOk"},{"Code": 301, "Descrip": "StatusMovedPermanently"},{"Code": 302, "Descrip": "StatusNotFound"},{"Code": 303, "Descrip": "StatusSeeOther"},{"Code": 307, "Descrip": "StatusTempporaryRedirect"},{"Code": 400, "Descrip": "StatusBadRequest"},{"Code": 401, "Descrip": "StatusUnauthorized"},{"Code": 402, "Descrip": "StatusPaymentRequired"},{"Code": 403, "Descrip": "StatusForbidden"},{"Code": 404, "Descrip": "StatusNotFound"},{"Code": 405, "Descrip": "StatusMethodNotAllowed"},{"Code": 418, "Descrip": "StatusTeaPot"},{"Code": 500, "Descrip": "StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range data {
		fmt.Println(v.Code, "-", v.Descrip)
	}

}
