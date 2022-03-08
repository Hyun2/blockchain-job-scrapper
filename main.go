package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var baseURL string = "https://www.jumpit.co.kr/api/positions?sort=relation&keyword=%EB%B8%94%EB%A1%9D%EC%B2%B4%EC%9D%B8"

type JumpitResult struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Result  struct {
		TotalCount int `json:"totalCount"`
		Page       int `json:"page"`
		Positions  []struct {
			ID               int      `json:"id"`
			JobCategory      string   `json:"jobCategory"`
			Logo             string   `json:"logo"`
			ImagePath        string   `json:"imagePath"`
			Title            string   `json:"title"`
			CompanyName      string   `json:"companyName"`
			TechStacks       []string `json:"techStacks"`
			ScrapCount       int      `json:"scrapCount"`
			ViewCount        int      `json:"viewCount"`
			Newcomer         bool     `json:"newcomer"`
			MinCareer        int      `json:"minCareer"`
			MaxCareer        int      `json:"maxCareer"`
			Locations        []string `json:"locations"`
			AlwaysOpen       bool     `json:"alwaysOpen"`
			ClosedAt         string   `json:"closedAt"`
			CompanyProfileID int      `json:"companyProfileId"`
			Celebration      int      `json:"celebration"`
			Scraped          bool     `json:"scraped"`
		} `json:"positions"`
	} `json:"result"`
}

func main() {
	pages := getPages()
	fmt.Println(pages)

	for i := 1; i <= pages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&page=" + strconv.Itoa(page)
	fmt.Println(pageURL)
}

func getPages() int {
	pages := 0

	res, err := http.Get(baseURL)
	checkErr(err)

	defer res.Body.Close()
	checkHttpResponse(res)

	body, _ := ioutil.ReadAll(res.Body) // response body is []byte

	var result JumpitResult
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	// fmt.Println(PrettyPrint(result))
	// fmt.Println(result.Result.TotalCount)

	if result.Result.TotalCount%16 == 0 {
		pages = result.Result.TotalCount / 16
	} else {
		pages = (result.Result.TotalCount / 16) + 1
	}

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkHttpResponse(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
