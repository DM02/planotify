package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/ecnepsnai/discord"
)

var globalElement string = ""

func main() {
	discord.WebhookURL = "https://discord.com/api/webhooks/950070131768311818/jQbnfY9pxvHH6EYDDce0u9bKA7BYuOWXpDDmHNaYCZp8hWk1aBFHbifxJuaOfSpJr1aQ"
	// Request the HTML page.
	res, err := http.Get("http://zsz1.edu.pl/plan/plany/o26.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	element := doc.Find("body > div > table > tbody > tr:nth-child(2) > td.op > table > tbody > tr > td:nth-child(1)")
	globalElement = element.Text()
	for {
		//http://localhost/Plan%20lekcji%20oddzia%c5%82u%20-%203pTP.html
		//http://zsz1.edu.pl/plan/plany/o26.html
		res, err := http.Get("http://zsz1.edu.pl/plan/plany/o26.html")
		if err != nil {
			log.Fatal(err)
		}
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		element := doc.Find("body > div > table > tbody > tr:nth-child(2) > td.op > table > tbody > tr > td:nth-child(1)")

		if globalElement == element.Text() {

			fmt.Println("Brak zmian", time.Now())

		} else {
			globalElement = element.Text()

			fmt.Println("Nastapila zmiana ", globalElement)
			fmt.Println("O godzinie: ", time.Now())
			discord.Say(fmt.Sprintf("Nastapila zmiana planu w dniu: %s", time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")))
			discord.Say("@everyone http://zsz1.edu.pl/plan/plany/o26.html")

		}
		time.Sleep(1 * time.Second)
	}

}
