package main

import (
	"fmt"

	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/antchfx/htmlquery"
)

func main() {
	cpuNumber := 3
	runtime.GOMAXPROCS(cpuNumber)
	var wait sync.WaitGroup
	wait.Add(cpuNumber)

	startTime := time.Now()

	go func() {
		for i := 0; i < 3; i++ {
			Crawler()
		}

		defer wait.Done()
	}()

	wait.Wait()
	elapsedTime := time.Since(startTime)

	fmt.Printf("실행시간: %s\n", elapsedTime)
}

func getHtml() (string, error) {

	url := "http://www.naver.com"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}

// Crawler is naver scraping
func Crawler() (string, error) {
	html, err := getHtml()

	if err != nil {
		return html, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(html))

	list := htmlquery.Find(doc, "//ul/li")

	for _, val := range list {
		fmt.Println("====== naver =======", htmlquery.InnerText(val))
	}

	return html, err
}
