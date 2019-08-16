package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randomAgent() string {
	b := make([]byte, rand.Intn(5)+20)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	var (
		FashionReps = "https://www.reddit.com/r/FashionReps/"
		GiftBags    = "search/?q=flair_name%3A\"GIFTBAG\"&restrict_sr=1&t=hour"
		UserName    = os.Getenv("USER")
		flagNoColor = flag.Bool("no-color", false, "Disable color output")
	)
	if *flagNoColor {
		color.NoColor = true
	}
	color.Set(color.FgCyan)
	logo := fmt.Sprintf("" +
		" _______ __  ___ __   _______             \n" +
		"|   _   |__.'  _|  |_|   _   .-----.-----.\n" +
		"|.  |___|  |   _|   _|.  1   |  _  |  _  |\n" +
		"|.  |   |__|__| |____|.  _   |_____|___  |\n" +
		"|:  1   |            |:  1    \\    |_____|\n" +
		"|::.. . |   v1.0.3   |::.. .  /           \n" +
		"`-------'            `-------'            \n")
	fmt.Println(logo)
	color.Set(color.FgWhite)
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", randomAgent())
		fmt.Println("[STATUS]: fetching giftbags")
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("[ERROR]:", err)
	})
	c.OnHTML("span", func(e *colly.HTMLElement) {
		e.ForEach("span", func(_ int, elem *colly.HTMLElement) {
			if strings.Contains(elem.Text, "GIFTBAG") {
				fmt.Println("[STATUS]: fetched a new giftbag", elem.Text)
				fmt.Println("[STATUS]: go unwrap your mf gifts", UserName)
				exec.Command("xdg-open", FashionReps+GiftBags).Start()
			} else {
				fmt.Println("[STATUS]: no giftbags at the moment")
			}
		})
	})
	c.Visit(FashionReps + GiftBags)
}
