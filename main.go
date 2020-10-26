package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	// import third party libraries
	"github.com/PuerkitoBio/goquery"
	//"github.com/chromedp/chromedp"
	//"github.com/chromedp/cdproto/dom"
	//"github.com/chromedp/cdptypes"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/meta", metaScrape)
    http.ListenAndServe(":8080", nil)
}


func metaScrape(w http.ResponseWriter, r *http.Request) {
	//pageURL := "https://www.w3schools.com/html/tryit.asp?filename=tryhtml_headings"
	//pageURL := "http://jonathanmh.com"
	//pageURL := "https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements" //links, headings
	//pageURL := "https://www.digikala.com/users/login-register/?_back=https://www.digikala.com/" //login
	//pageURL := "https://www.zanbil.ir/login/auth?forwardUri=%2F" //login
	pageURL := "https://github.com/login"
	//pageURL := "https://okala.com/account/login" //login action
	// pageURL := "http://metalsucks.net" //inaccible code 403
	//error handing
	res, err := http.Get(pageURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("*****************************")
	fmt.Println(res.Body)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	//doc, err := goquery.NewDocument(pageURL)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {

		log.Fatal(err)
	}

	var metaDescription string
	var pageTitle string

	// use CSS selector found with the browser inspector
	// for each, use index and item
	pageTitle = doc.Find("title").Contents().Text()

	doc.Find("meta").Each(func(index int, item *goquery.Selection) {
		if (item.AttrOr("name", "") == "description") {
			metaDescription = item.AttrOr("content", "")
		}

	})
	fmt.Printf("Page Title: '%s'\n", pageTitle)
	fmt.Printf("Meta Description: '%s'\n", metaDescription)

	fmt.Println("================heading")
	//find heading : six level
	doc.Find("h1").Each(func(i int, h1 *goquery.Selection) {
		fmt.Println(i)
	})

	fmt.Println("================link+++++++++++++++++++++++")
	//internal and external links
	//doc.Find("body a").Each(func(index int, item *goquery.Selection) {
	//    linkTag := item
	//    link, _ := linkTag.Attr("href")
	//    linkText := linkTag.Text()
	//checkExtrernal := strings.HasPrefix(link, "http")
	//external := "true"
	//domain := ""
	//if checkExtrernal {
	//	external = "true"
	//
	//} else {
	//	external = "false"
	//	domain = r.URL.Host
	//	fmt.Println(domain)
	//	return
	//
	//}
	//
	//
	//
	//
	//res, err := http.Get(link)
	//    if err != nil {
	//        log.Fatal(err)
	//    }
	//    defer res.Body.Close()
	//
	//    fmt.Printf("Link #%d: '%s' - '%s'-  '%s', '%s', '%s' \n", index, linkText, link,external, strings.HasPrefix(link, "http"), res.StatusCode )
	//})

	//login
	doc.Find("html body").Each(func(_ int, item *goquery.Selection) {
		// for debug.
		println(item.Size()) // return 1

		// if len(s.Nodes) > 0 && s.Nodes[0].Type == html.ElementNode {
		//   println(s.Nodes[0].Data)
		//}
		if (item.AttrOr("name", "") == "password") {
			fmt.Println("here is login form")
		}
	})

	fmt.Println("--------------------------------------------------------------------------")
	//doc.Find("meta").Each(func(i int, s *goquery.Selection) {
	//	description, _ := s.Attr("content")
	//	//a: = s.HasNodes("DOCTYPE")
	//	//fmt.Println("************************************")
	//	//fmt.Println(s.Text())
	//
	//	fmt.Printf("Description field: %s\n", description)
	//})
	//
	//doc.Find("DOCTYPE").Each(func(i int, s *goquery.Selection) {
	//	fmt.Println(s.Text())
	//
	//
	//})
	//doc.Each(func(_ int, item *goquery.Selection) {
	//	fmt.Println(">>>>>>>>>>>>..........")
	//	fmt.Println(item)
	//	fmt.Println(goquery.NodeName(item))
	//	fmt.Println(item.Text())
	//})

	//doc.Find("DOCTYPE").Each(func(i int, s *goquery.Selection) {
	//	fmt.Println(s.Text())
	//})

	//fmt.Println("..................")
	//fmt.Println(doc.Text())

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>..................<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

	password, ok := doc.Find("input").Attr("id")
	fmt.Println(password)
	fmt.Println(ok)
	if !ok {

	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>11111111111111111111<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	 login, ok := doc.Find("form").Attr("id")
	fmt.Println(login)
	fmt.Println(ok)
	if !ok {

	}

	checkLogin := strings.Contains(login, "login")
	if checkLogin {

	}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>=======")
	doc.Find("html body").Each(func(_ int, item *goquery.Selection) {
		// for debug.
		println(item.Size()) // return 1

		// if len(s.Nodes) > 0 && s.Nodes[0].Type == html.ElementNode {
		//   println(s.Nodes[0].Data)
		//}
		fmt.Println(item.AttrOr("name", ""))
		if (item.AttrOr("name", "") == "password") {
			fmt.Println("here is login form")
		}
	})

	token, ok := doc.Find("input[name='password']").Attr("id")
	fmt.Println(token)
	fmt.Println(ok)

	//doc.Find(".fheader").Each(func(i int, s *goquery.Selection) {
	//	//name := strings.TrimSpace(s.Text())
	//	fmt.Println("==========================.........||||||||||||||||")
	//	fmt.Println(s.Find("input").Attr("id"))
	//	//project := Project{
	//	//	Name: name,
	//	//}
	//	//
	//	//projects = append(projects, project)
	//})




	doc.Find("body input").Each(func(index int, item *goquery.Selection) {

	   itemId, _ := item.Attr("id")



	   fmt.Printf("Link #%d: '%s' -  \n", index ,itemId)
	})


	fmt.Println(res.Body)
	b, err := html.Parse(res.Body)
	if err != nil {
		 fmt.Errorf("Cannot parse page")
	}
	fmt.Println(b)
	fmt.Println(b.Data)
	fmt.Println(b.Namespace)

	//html, err := dom.GetOuterHTML().WithNodeID(cdptypes.NodeID(0)).Do(ctxt, c)
	//fmt.Println(html)
	//fmt.Println(err)
	fmt.Println("222222222222222222222222222222222")
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("errrr")
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)

}


//func findHeading(doc) {


//}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
