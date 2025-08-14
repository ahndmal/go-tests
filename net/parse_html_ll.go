package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getById(id string, htm string) string {
	doc, err := html.Parse(strings.NewReader(htm))
	if err != nil {
		panic(err)
	}

	var elem string

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" {
					if attr.Val == id {
						elem = attr.Val
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	return elem
}

func getWeatherHtml() string {
	resp, err := http.Get("https://www.timeanddate.com/sun/ukraine/sumy")
	if err != nil {
		fmt.Printf("Error perf GET: %v", err)
	}

	bts, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error reading bytes from response : %v", err)
	}

	return string(bts)
}

func main() {
	htmlData := getWeatherHtml()

	doc, err := html.Parse(strings.NewReader(htmlData))
	if err != nil {
		panic(err)
	}

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			for _, attr := range n.Attr {
				if attr.Key == "id" {
					if attr.Val == "as-monthsun" {
						tbody := n.FirstChild.NextSibling

						var trElem *html.Node
						trElem = tbody.FirstChild
						fmt.Println(trElem.Data)

						trElem = trElem.NextSibling
						if trElem != nil {
							for trElem != nil {

								th := trElem.FirstChild
								dayNum := th.FirstChild.Data

								// get TD
								tdOne := th.NextSibling

								sunrise := tdOne.FirstChild.Data
								sunset := tdOne.NextSibling.FirstChild.Data
								dayLength := tdOne.NextSibling.NextSibling.FirstChild.Data
								dayDiff := tdOne.NextSibling.NextSibling.NextSibling

								// astronomical twilight
								astroStart := dayDiff.NextSibling
								astroEnd := astroStart.NextSibling

								// Nautical twilight
								nautStart := astroEnd.NextSibling
								nautEnd := nautStart.NextSibling

								// Civil Twilight
								civilStart := nautEnd.NextSibling
								civilEnd := civilStart.NextSibling

								// Solar noon
								solarTime := civilEnd.NextSibling
								solarMilKm := solarTime.NextSibling

								// print
								fmt.Printf("%s | %s | %s | %s | %s | %s | %s | %s  \r\n", dayNum, sunrise, sunset, dayLength, dayDiff.FirstChild.Data, astroStart.FirstChild.Data, astroEnd.FirstChild.Data, solarMilKm.FirstChild.Data)

								// next TR
								trElem = trElem.NextSibling
							}
						}
					}
				}

			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)
}
