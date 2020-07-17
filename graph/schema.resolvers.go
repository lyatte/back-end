package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"back_end/graph/generated"
	"back_end/graph/model"
	"context"
	"errors"
	"fmt"
	"log"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input *model.NewVideo) (*model.Video, error) {
	video := model.Video{
		VideoTitle:       input.VideoTitle,
		VideoDescription: input.VideoDescription,
		VideoCategory:    input.VideoCategory,
		VideoLike:        input.VideoLike,
		VideoDislike:     input.VideoDislike,
		VideoPrivacy:     input.VideoPrivacy,
		VideoPremium:     input.VideoPremium,
		VideoRestriction: input.VideoRestriction,
		VideoThumbnail:   input.VideoThumbnail,
		Video:            input.Video,
		ChannelID:        input.ChannelID,
		VideoViews:       input.VideoViews,
		VideoRegion:      input.VideoRegion,
		Day:              input.Day,
		Month:            input.Month,
		Year:             input.Year,
	}

	_, err := r.DB.Model(&video).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert failed!")
	} else {
		log.Println("Insert success!")
		return &video, nil
	}
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, videoID string, input *model.NewVideo) (*model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, videoID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetVideo(ctx context.Context) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Order("video_id").Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed.")
	}

	return videos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
