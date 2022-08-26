package api

import (
	"context"
	"errors"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Order string

type CommentList struct {
	VideoId     string
	Order       Order
	MaxComments int
}

func (api *Api) GetComments(list CommentList) ([]string, error) {
	var comments []string

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	service, err := youtube.NewService(ctx, option.WithAPIKey(api.APIKey))
	if err != nil {
		cancel()
		return nil, err
	}

	commentThreadsService := youtube.NewCommentThreadsService(service)

	call := commentThreadsService.List([]string{"snippet"})
	call.TextFormat("plainText")
	call.VideoId(list.VideoId)
	call.Order(string(list.Order))
	call.MaxResults(100)

	err = call.Pages(ctx, func(page *youtube.CommentThreadListResponse) error {
		for _, item := range page.Items {
			comment := item.Snippet.TopLevelComment.Snippet.TextDisplay
			comments = append(comments, comment)

			if len(comments) == list.MaxComments {
				cancel()
				break
			}
		}
		return nil
	})

	if err != nil && !errors.Is(err, context.Canceled) {
		cancel()
		return nil, err
	}
	return comments, nil
}
