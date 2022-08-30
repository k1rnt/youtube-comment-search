package main

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/k1rnt/youtube-comment-search/pkg/api"
	"github.com/k1rnt/youtube-comment-search/pkg/search"
)

var (
	videoId    = flag.String("video_id", "UKZt1vq8bKI", "video id")
	maxResults = flag.Int("max_results", 100, "max comments")
	keyword    = flag.String("keyword", "", "keyword")
	regex      = flag.String("regex", "", "regex")
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	youtubeApi := api.NewApi()
	comments, err := youtubeApi.GetComments(api.CommentList{
		VideoId:     *videoId,
		Order:       api.Order("time"),
		MaxComments: *maxResults,
	})
	if err != nil {
		panic(err)
	}
	var res []string
	for _, comment := range search.Keyword(comments, *keyword) {
		res = append(res, comment)
	}
	for _, comment := range search.Regex(res, *regex) {
		res = append(res, comment)
	}
	for i, comment := range res {
		fmt.Printf("%d: %s\n", i, comment)
	}
}
