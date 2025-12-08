package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)





func StringSetToSlice (set map[string]struct{}) ([]string){
	resultSlice := make([]string, len(set))
	for key := range set{
		resultSlice = append(resultSlice, key)
	}
	return resultSlice
}


func main(){
	godotenv.Load("../.env")
	gamesCodesURL := os.Getenv("GAMECODES_URL")
	redemptionServiceURL := os.Getenv("REDEMPTION_SERVICE_URL")
	addTicketsEndpoint := redemptionServiceURL + "/v1/coupons"
	codes, err := ScrapCodesFromURL(gamesCodesURL)
	if err != nil{
		log.Fatalf("Couldn't scrap HTML from URL, got error: %v", err.Error())
	}
	// TODO inject actual backend url
	err = SendCurrentCoupons(addTicketsEndpoint,codes)
	if err != nil{
		log.Printf("Got error when trying to send coupon codes: %v", err.Error())
	}	
}



func ScrapCodesFromURL(url string) (map[string](struct{}), error) {
	req, err := http.NewRequest("GET",url,nil)
	codes := make(map[string](struct{}),10)
	if err != nil{
		return map[string](struct{}){}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil{
		return map[string](struct{}){}, err
	}

	defer res.Body.Close()
	if err != nil{
		return map[string](struct{}){}, err
	}	
	err = ExtractCodesFromBody(res.Body, &codes)
	if err != nil{
		return map[string](struct{}){}, err
	}
	return codes, nil
}

func ExtractCodesFromBody(body io.Reader, codes *map[string](struct{})) (error){
	doc, err := goquery.NewDocumentFromReader(body)
    if err != nil {
        return err
    }
	doc.Find("body").Find("td").Find("a").Each(func(i int, s *goquery.Selection) {
		elem, exists := s.Attr("data-gamecode")
		if exists{
			(*codes)[elem] = struct{}{}
		}	
	})
	fmt.Println(*codes)
	return nil
}

func SendCurrentCoupons(url string, extractedCoupons map[string](struct{})) (error){
	type bodyParams struct{
		Codes []string `json:"coupon_codes"`
	}
	payloadData := bodyParams{Codes: StringSetToSlice(extractedCoupons)}
	log.Println("Trying to send the following fetched codes: ",payloadData)
	marshalledPayload, err := json.Marshal(payloadData)
	if err != nil{
		return err
	}

	req, err := http.NewRequest("POST",url, bytes.NewBuffer(marshalledPayload))
	if err != nil{
		return err
	}
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return err
	}
	if resp.StatusCode != http.StatusCreated{
		return fmt.Errorf("post gamecodes did not result in 200 http code, got code: %w", resp.StatusCode)
	}
	return nil
}