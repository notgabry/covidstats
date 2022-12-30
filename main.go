package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type Covid struct {
	Region    string `json:"denominazione_regione"`
	Positives int    `json:"totale_positivi"`
	Isolation int    `json:"isolamento_domiciliare"`
}

func main() {
	baseURL := "https://raw.githubusercontent.com/pcm-dpc/COVID-19/master/dati-json/dpc-covid19-ita-regioni-latest.json"
	resp, err := http.Get(baseURL)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	var CovidStruct []Covid
	json.Unmarshal(bodyBytes, &CovidStruct)

	var positives int
	for _, a := range CovidStruct {
		positives += a.Positives
	}
	var isolation int
	for _, a := range CovidStruct {
		isolation += a.Isolation
	}
	fmt.Printf(`
%v the latest cases of %s
|
|_ %v %s
|
|_ %v %s
`, color.HiGreenString("Check out"), color.HiBlueString("Covid-19 in Italy"), color.HiWhiteString("Positives"), color.YellowString(fmt.Sprint(Comma(int64(positives)))), color.HiWhiteString("Isolated"), color.YellowString(fmt.Sprint(Comma(int64(isolation)))))
}

func Comma(value int64) string {
	formattedNumber := fmt.Sprintf("%d", value)

	for i := len(formattedNumber) - 3; i > 0; i -= 3 {
		formattedNumber = fmt.Sprintf("%v,%v", formattedNumber[:i], formattedNumber[i:])
	}

	return formattedNumber
}
