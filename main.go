package main

import(
	"fmt"
	"net/http"
	"net/url"
	"log"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func main(){
	u, err := url.ParseRequestURI(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	if u.Host != "pastebin.com" {
		fmt.Println("not a pastebin URL fucktard")
		os.Exit(1)
	}

	if u.IsAbs() != true {
		fmt.Println("something went wrong lol, ask Schmenn for help")
		os.Exit(1)
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		fmt.Println("bruh, at least at \"http://\" or \"https://\" to the URL")
		os.Exit(1)
	}

	/*var output string

	for i, a := range os.Args{
		if a == "-o" {
			output = os.Args[i+1]
		} else {
			output = u.
		}
	}*/

	res, err := http.Get(u.String())
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalln("status code error: " + string(res.StatusCode) + string(res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	text := ""

	doc.Find("div.wrap div.container div.content div.post-view textarea.textarea").Each(func(i int, s *goquery.Selection) {
		text += s.Text()
	})

	fmt.Println(text)
}


