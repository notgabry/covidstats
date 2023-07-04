package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type Covid struct {
	Region    string `json:"denominazione_regione"`
	Positives int    `json:"totale_positivi"`
	Isolation int    `json:"isolamento_domiciliare"`
}

func main() {
	baseURL := "https://raw.githubusercontent.com/pcm-dpc/COVID-19/master/dati-json/dpc-covid19-ita-regioni-latest.json"
	res, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("Impossible to reach github servers.")
		os.Exit(1)
	}

	defer res.Body.Close()
	BodyBytes, _ := io.ReadAll(res.Body)

	var CovidStruct []Covid
	json.Unmarshal(BodyBytes, &CovidStruct)

	var positives int
	var isolation int
	for _, a := range CovidStruct {
		positives += a.Positives
		isolation += a.Isolation
	}

	light := color.New(color.FgHiBlue).SprintFunc()
	yellow := color.New(color.FgHiYellow).SprintFunc()

	fmt.Printf("\n%s\n%s Positives\n%s Isolated\n", light(time.Now().Format("Monday 2 January 2006")), yellow(Comma(positives)), yellow(Comma(isolation)))
}

func Comma(value int) string {
	FormattedNumber := fmt.Sprintf("%d", value)

	for i := len(FormattedNumber) - 3; i > 0; i -= 3 {
		FormattedNumber = fmt.Sprintf("%s,%s", FormattedNumber[:i], FormattedNumber[i:])
	}

	return FormattedNumber
}
