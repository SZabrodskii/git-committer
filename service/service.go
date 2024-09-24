package service

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func GetRandomAnekdot() (string, error) {
	resp, err := http.Get("https://www.anekdot.ru/random/anekdot/")
	if err != nil {
		return "", fmt.Errorf("failed to fetch joke: %w", err)
	}
	defer resp.Body.Close()

	doc := html.NewTokenizer(resp.Body)
	var joke string
	found := false
	for {
		tt := doc.Next()
		switch tt {
		case html.ErrorToken:
			if joke == "" {
				return "", fmt.Errorf("failed to parse joke")
			}
			return joke, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			t := doc.Token()
			if t.Data == "div" {
				for _, attr := range t.Attr {
					if attr.Key == "class" && attr.Val == "text" {
						found = true
					}
				}
			}
		case html.TextToken:
			if found {
				joke = strings.TrimSpace(doc.Token().Data)
				found = false
			}
		}
	}
}
