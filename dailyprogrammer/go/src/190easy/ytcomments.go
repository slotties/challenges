package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
)

/*
	Specific part of Challenge:
	Download comments of a youtube video.
*/

func LoadComments(url string) []string {
	rsp, err := http.Get("https://plus.googleapis.com/u/0/_/widget/render/comments?first_party_property=YOUTUBE&href=" + url)
	if err == nil {
		defer rsp.Body.Close()
		html, err := ioutil.ReadAll(rsp.Body)
		if err == nil {
			return parseComments(string(html))
		}
	}

	return make([]string, 0)
}

func parseComments(html string) []string {
	re, _ := regexp.Compile("(?s)<div class=\"Ct\">(.*?)</div>")
	res := re.FindAllStringSubmatch(html, -1)

	comments := make([]string, len(res))
	for idx, el := range res {
		comments[idx] = el[1]
	}

	return comments
}
