package jobscrapper

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=%EB%B8%94%EB%A1%9D%EC%B2%B4%EC%9D%B8"

func main() {
	pages := getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)

	defer res.Body.Close()
	checkHttpResponse(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination-list")

	return 0
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
