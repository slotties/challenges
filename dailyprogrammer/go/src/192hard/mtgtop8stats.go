package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"errors"
	"regexp"
	"strconv"
	//"fmt"
)

type extractor func([]string)

func GetMajorEventCardStatistics(format string) (map[string]int, error) {
	events, err := LoadLatestMajorEvents(format)
	if err != nil {
		return nil, err
	}

	cards := make(map[string]int)
	for eventId, _ := range events {		
		decks, err := LoadEventDecks(eventId, format)
		if err == nil {
			for deckId, _ := range decks {
				deckCards, err := LoadCards(eventId, format, deckId)
				if err == nil {
					for card, amount := range deckCards {
						cards[card] += amount
					}
				}
			}
		}
	}

	return cards, nil
}

func LoadCards(eventId string, format string, deckId string) (map[string]int, error) {
	html, err := loadHtml("http://www.mtgtop8.com/event?e=" + eventId + "&d=" + deckId + "&f=" + format)
	if err != nil {
		return nil, err
	}

	cards := make(map[string]int)
	extractor := func(match []string) {
		name := strings.TrimSpace(match[2])
		amount, err := strconv.Atoi(match[1])
		if err != nil {
			amount = 0
		}
		cards[name] = amount
	}

	err = extractData(html, "<table border=0 class=Stable width=98%>", "<div align=center>", "(?U)([0-9]+) <span .*>(.+)</span>", extractor)
    return cards, err
}

func LoadEventDecks(eventId string, format string) (map[string]string, error) {
	html, err := loadHtml("http://www.mtgtop8.com/event?e=" + eventId + "&f=" + format)
	if err != nil {
		return nil, err
	}

	decks := make(map[string]string)
	extractor := func(match []string) {
		decks[match[1]] = strings.TrimSpace(match[2])
	}

	err = extractData(html, "", "", "<a .*href=event\\?.*d=(.*)&.*>(.*)</a>", extractor)
    return decks, err
}

func LoadLatestMajorEvents(format string) (map[string]string, error) {
	html, err := loadHtml("http://www.mtgtop8.com/format?f=" + format)
	if err != nil {
		return nil, err
	}

	events := make(map[string]string)
	extractor := func(match []string) {
		events[match[1]] = strings.TrimSpace(match[2])
	}	
	err = extractData(html, "Last major events", "</table>", "<a href=event\\?e=(.*)&.*>(.*)</a>", extractor)

	return events, err
}

func extractData(text string, startStr string, endStr string, expression string, extractorFn extractor) error {
	if startStr != "" && endStr != "" {
		var err error
		text, err = extractPart(text, startStr, endStr)
		if err != nil {
			return err
		}
	}

	re := regexp.MustCompile(expression)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		extractorFn(match)
	}

	return nil
}

func extractPart(text string, start string, end string) (string, error) {
	startIdx := strings.Index(text, start)
	if startIdx >= 0 {
		text = text[startIdx:]
		endIdx := strings.Index(text, end)
		if endIdx > 0 {
			return text[:endIdx], nil
		} else {
			return "", errors.New("Could not find '" + end + "'' in text")
		}
	} else {
		return "", errors.New("Could not find '" + start + "'' in text")
	}
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