package main

import (
	"fmt"
)

/*
	Challenge:
	Your task is to scrape N (You decide but generally, the higher the sample, the more accurate) number of comments from a YouTube video of your choice and then analyse their sentiments based on a short list of happy/sad keywords
	Analysis will be done by seeing how many Happy/Sad keywords are in each comment. If a comment contains more sad keywords than happy, then it can be deemed sad.
	Here"s a basic list of keywords for you to test against. I"ve ommited expletives to please all readers...
	happy = ["love","loved","like","liked","awesome","amazing","good","great","excellent"]
	sad = ["hate","hated","dislike","disliked","awful","terrible","bad","painful","worst"]
	Feel free to share a bigger list of keywords if you find one. A larger one would be much appreciated if you can find one.

	http://www.reddit.com/r/dailyprogrammer/comments/2nauiv/20141124_challenge_190_easy_webscraping_sentiments/
*/
func main() {
	// FIXME: get through parameter
	url := "https://www.youtube.com/watch?v=BOByH_iOn88"
	analyzer := New()
	// Happy words
	analyzer.Add("love", 1)
	analyzer.Add("loved", 1)
	analyzer.Add("like", 1)
	analyzer.Add("liked", 1)
	analyzer.Add("awesome", 1)
	analyzer.Add("amazing", 1)
	analyzer.Add("good", 1)
	analyzer.Add("gread", 1)
	analyzer.Add("excellent", 1)
	// Sad words
	analyzer.Add("hate", -1)
	analyzer.Add("hated", -1)
	analyzer.Add("dislike", -1)
	analyzer.Add("disliked", -1)
	analyzer.Add("awful", -1)
	analyzer.Add("terrible", -1)
	analyzer.Add("bad", -1)
	analyzer.Add("painful", -1)
	analyzer.Add("worst", -1)

	comments := LoadComments(url)
	for _, comment := range comments {
		score := analyzer.Score(comment)
		fmt.Printf("This sentence has a score of %d\n", score)
	}
}