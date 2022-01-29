package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//---------------- Using Goroutines and channels, testing the speed for simultaneous 4 Get requests ----------------

var startProgram = time.Now()

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		time := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, time}

		resp.Body.Close()
	}
}

func testGETspeed() {

	listOfUrls := []string{"https://www.google.com", "https://www.youtube.com/", "https://www.facebook.com/", "https://www.amazon.com/"}
	results := make(chan result)
	for _, url := range listOfUrls {
		go get(url, results)
	}
	for range listOfUrls {
		r := <-results
		if r.err != nil {
			log.Printf("%v %v", r.url, r.err)
		} else {
			log.Printf("%v %v", r.url, r.latency)
		}

	}
	duration := time.Since(startProgram).Round(time.Millisecond)
	fmt.Println(duration)

}

// ----------------------------------------------End of speedTest program----------------------------------------------

func main() {
	testGETspeed()
	// Making my own http client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	// ------------ HTTP clients use a default Request type which we can also customize ------------

	// Creating and modifying HTTP request
	req, err := http.NewRequest("GET", "https://www.devdungeon.com/", nil)
	if err != nil {
		log.Fatal(err)
	}
	// Adding a user-agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html")
	// Creating our own cookie
	mycookie := &http.Cookie{
		Name:  "mayhul",
		Value: "jindal",
	}
	req.AddCookie(mycookie)
	// Displays the cookie we created and shows us the header where the cookie is attached
	fmt.Println(req.Cookies())
	fmt.Println(req.Header)

	// ------------ Done creating our own http get request ------------

	// Make HTTP GET request
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	// Making a file to store the output
	output_file, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer output_file.Close()
	// coping response body to the output_file
	_, err = io.Copy(output_file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Searching strings techniques
	// ------------ Using sub-string matching to find the element ------------

	output, err := os.Open("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	output_in_bytes, err := ioutil.ReadAll(output)
	if err != nil {
		log.Fatal(err)
	}
	output_in_string := string(output_in_bytes)
	startindex := strings.Index(output_in_string, "<title>")
	if startindex == -1 {
		fmt.Println("No such string found ")
		os.Exit(0)
	}
	// The start index of the title is the index of the first
	// character, the < symbol. We don't want to include
	// <title> as part of the final value, so let's offset
	// the index by the number of characers in <title>
	startindex += 7
	endindex := strings.Index(output_in_string, "</title>")
	if endindex == -1 {
		fmt.Println("no such string found")
		os.Exit(0)
	}
	searched_string := output_in_string[startindex:endindex]
	fmt.Println(searched_string)
	// ------------ Using regular expressions ------------

	re := regexp.MustCompile("<!--(.|\n)*?-->")
	comments := re.FindAllString(string(output_in_string), -1)
	if comments == nil {
		fmt.Println("No matches.")
	} else {
		for _, comment := range comments {
			fmt.Println(comment)
		}
	}
	// ------------ Using goquery ------------

	// first making a function processTags which will run for all the the tags found
	html_code, err := os.Open("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer html_code.Close()
	document, err := goquery.NewDocumentFromReader(html_code)
	if err != nil {
		log.Fatal(err)
	}

	document.Find("a").Each(processTags)
	// Parse a complex url
	host_file, err := os.Create("host_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer host_file.Close()
	for _, hostname := range hostnameslice {
		_, err = host_file.WriteString(hostname + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	// ------------ Making a login post request ------------
	login()
	// ------------ X ------------------------- X --------------
	fmt.Println("######## SUCCESS #######")
}

// ------------ Creating processTag function as mentioned above ------------
var hostnameslice []string

func processTags(index int, element *goquery.Selection) {
	// see if we get anything in the attribute of the tag
	attribute, exist := element.Attr("href")
	if exist && strings.Contains(attribute, "http") {
		parsedurl, err := url.Parse(attribute)
		if err != nil {
			log.Fatal(err)
		}
		hostnameslice = append(hostnameslice, parsedurl.Host)
		fmt.Println(attribute)
	} else {
		fmt.Println("")
	}

}

// To upload a file, use Post instead of PostForm, provide
// a content type like application/json or application/octet-stream,
// and then provide the an io.Reader with the data

// http.Post("http://example.com/upload", "image/jpeg", &buff)

// ------------X------------X------------
