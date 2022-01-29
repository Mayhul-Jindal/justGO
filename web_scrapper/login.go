package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func login() {
	jar, _ := cookiejar.New(nil)

	app := App{
		Client: &http.Client{Jar: jar},
	}
	app.getToken()
	app.post()
	app.checkingLogin()
}

const loginurl = "http://quotes.toscrape.com/login"

var token string

type App struct {
	Client *http.Client
}

func (app *App) getToken() {
	client := app.Client
	res, err := client.Get(loginurl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	document.Find("input").Each(func(index int, element *goquery.Selection) {
		attribute_value, exsist := element.Attr("value")
		if exsist {
			attribute_name, exsist := element.Attr("name")
			if exsist && attribute_name == "csrf_token" {
				token = attribute_value
				fmt.Println(token)
			}
		}

	})

}
func (app *App) post() {
	client := app.Client
	// ------Sample payload------
	// csrf_token: "dLGFBjIDEZWXyqsifAQYpuwrVKnzTcRlkvOSetCMNhPgxUHbJoma",
	// username: "mayhul",
	// password: "jindal",
	data := url.Values{
		"csrf_token": {token},
		"username":   {"mayhul"},
		"password":   {"passwd"},
	}
	fmt.Println(data)
	res, err := client.PostForm(loginurl, data)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

func (app *App) checkingLogin() {
	client := app.Client

	res, err := client.Get("http://quotes.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		attribute_value, exsist := element.Attr("href")
		if exsist && attribute_value == "/logout" {
			fmt.Println("Login successful")
		}
	})
}
