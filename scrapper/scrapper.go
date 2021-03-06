package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

// Scrape indeed page
func Scrape(term string) {
	var baseURL = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalpages := getPages(baseURL)
	fmt.Println(totalpages)
	for i := 0; i < totalpages; i++ {
		go getPage(i, baseURL, c)
	}
	for i := 0; i < totalpages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Extracted jobs ", len(jobs))
}
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	//	c := make(chan []string)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	//	for _, job := range jobs {
	//	go writeJob(job, c)
	//}
	//	for i := 0; i < len(jobs); i++ {
	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		// jobSlice := <-c
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

// func writeJob(job extractedJob, c chan<- []string) {
// 	jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
// 	c <- jobSlice
// }

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkRes(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCard := doc.Find(".jobsearch-SerpJobCard")
	searchCard.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i := 0; i < searchCard.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}
func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkRes(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, card *goquery.Selection) {
		pages = card.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkRes(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed statuscode :", res.StatusCode)
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
