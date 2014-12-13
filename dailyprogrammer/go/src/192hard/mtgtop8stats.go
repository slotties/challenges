package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	//"errors"
	//"fmt"
	"regexp"
	"strconv"
)

func LoadCards(eventId string, format string, deckId string) (map[string]int, error) {
	html, err := loadHtml("http://www.mtgtop8.com/event?e=" + eventId + "&d=" + deckId + "&f=" + format)
	if err != nil {
		return nil, err
	}

	return parseCards(html)
}

func parseCards(html string) (map[string]int, error) {
	html = extractPart(html, "<table border=0 class=Stable width=98%>", "<div align=center>")

	cards := make(map[string]int)
	re := regexp.MustCompile("(?U)([0-9]+) <span .*>(.+)</span>")
	matches := re.FindAllStringSubmatch(html, -1)

	for _, match := range matches {
		name := strings.TrimSpace(match[2])
		amount, err := strconv.Atoi(match[1])
		if err != nil {
			amount = 0
		}
		cards[name] = amount
	}
    
    return cards, nil	
}

func LoadEventDecks(eventId string, format string) (map[string]string, error) {
	html, err := loadHtml("http://www.mtgtop8.com/event?e=" + eventId + "&f=" + format)
	if err != nil {
		return nil, err
	}

	return parseEventDecks(html)
}

func parseEventDecks(html string) (map[string]string, error) {
	decks := make(map[string]string)
	re := regexp.MustCompile("<a .*href=event\\?.*d=(.*)&.*>(.*)</a>")
	matches := re.FindAllStringSubmatch(html, -1)

	for _, match := range matches {
		decks[match[1]] = strings.TrimSpace(match[2])
	}
    
    return decks, nil

}

func LoadLatestMajorEvents(format string) (map[string]string, error) {
	html, err := loadHtml("http://www.mtgtop8.com/format?f=" + format)
	if err != nil {
		return nil, err
	}

	return parseMajorEvents(html)
}

func parseMajorEvents(html string) (map[string]string, error) {
	html = extractPart(html, "Last major events", "</table>")

	events := make(map[string]string)
	re := regexp.MustCompile("<a href=event\\?e=(.*)&.*>(.*)</a>")
	matches := re.FindAllStringSubmatch(html, -1)

	for _, match := range matches {
		events[match[1]] = strings.TrimSpace(match[2])
	}
    
    return events, nil
}

func extractPart(text string, start string, end string) string {
	startIdx := strings.Index(text, start)
	if startIdx >= 0 {
		text = text[startIdx:]
		endIdx := strings.Index(text, end)
		if endIdx > 0 {
			return text[:endIdx]
		}
	}

	return ""
}

func loadHtml(url string) (string, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	} else {
		return string(body), nil
	}
}