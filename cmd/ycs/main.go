package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/k1rnt/youtube-comment-search/pkg/api"
)

var (
	videoId    = flag.String("video_id", "UKZt1vq8bKI", "video id")
	maxResults = flag.Int("max_results", 100, "max comments")
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
	for _, comment := range comments {
		println(comment)
	}
}
