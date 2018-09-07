package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var soulapi SoulApi
	structAPI := &Soul{}
	soulapi = structAPI

	doc, err := soulapi.get("https://www.feixiaohao.com/")
	if err != nil {
		log.Fatal(err)
	}

	allData := []CoinInfo{}
	doc.Find("table tbody tr").Each(
		func(i int, selector *goquery.Selection) {
			var coinInfo CoinInfo
			Rank := selector.Find("td").Eq(0).Text()
			CoinName := strings.TrimSpace(selector.Find("td").Eq(1).Text())
			CurrentCount := selector.Find("td").Eq(2).Text()
			CurrentPrice := selector.Find("td").Eq(3).Text()
			CurrentMark := selector.Find("td").Eq(4).Text()
			Count := selector.Find("td").Eq(5).Text()
			Change := strings.TrimSpace(selector.Find("td").Eq(6).Text())
			coinInfo = CoinInfo{
				Rank:         Rank,
				Name:         CoinName,
				CurrentCount: CurrentCount,
				CurrentPrice: CurrentPrice,
				CurrentMark:  CurrentMark,

				Count:  Count,
				Change: Change,
			}
			allData = append(allData, coinInfo)
		})

	aimString, err := json.MarshalIndent(allData, "", " ")
	fmt.Println(string(aimString))

}

type SoulApi interface {
	get(string) (*goquery.Document, error)

	post(string, map[string]interface{}) ([]byte, error)

	put(string, map[string]interface{}) ([]byte, error)

	delete(string) (int, error)
}

type Soul struct {
	myHost     string
	myLocation string
}
type CoinInfo struct {
	Rank         string `json:"rank"`
	Name         string `json:"name"`
	CurrentCount string `json:"current_count"`
	CurrentPrice string `json:"current_price"`
	CurrentMark  string `json:"current_mark"`
	Count        string `json:"count"`
	Change       string `json:"change"`
}

func (this *Soul) get(url string) (*goquery.Document, error) {

	response, err := goquery.NewDocument(url)
	return response, err
}
func (s Soul) post(url string, params map[string]interface{}) ([]byte, error) {
	return nil, nil
}
func (s Soul) put(url string, params map[string]interface{}) ([]byte, error) {
	return nil, nil
}
func (s Soul) delete(url string) (int, error) {
	return 0, nil
}
