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
		ChannelName: input.ChannelName,
		ChannelIcon: input.ChannelIcon,
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

func (r *mutationResolver) CreateChannel(ctx context.Context, input *model.NewChannel) (*model.Channel, error) {
	channel := model.Channel{
		ChannelName:          input.ChannelName,
		ChannelBackground:    input.ChannelBackground,
		ChannelIcon:          input.ChannelIcon,
		ChannelSubscribers:   input.ChannelSubscribers,
		ChannelDescription:   input.ChannelDescription,
		ChannelJoinDateDay:   input.ChannelJoinDateDay,
		ChannelJoinDateMonth: input.ChannelJoinDateMonth,
		ChannelJoinDateYear:  input.ChannelJoinDateYear,
	}

	_, err := r.DB.Model(&channel).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert failed!")
	} else {
		log.Println("Insert success!")
		return &channel, nil
	}
}

func (r *mutationResolver) UpdateChannel(ctx context.Context, channelID string, input *model.NewChannel) (*model.Channel, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteChannel(ctx context.Context, channelID string) (bool, error) {
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

func (r *queryResolver) GetVideoByID(ctx context.Context, videoID int) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("video_id = ?", videoID).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	return &video, nil
}

func (r *queryResolver) GetChannel(ctx context.Context) ([]*model.Channel, error) {
	var channels []*model.Channel

	err := r.DB.Model(&channels).Order("channel_id").Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed.")
	}

	return channels, nil
}

func (r *queryResolver) GetChannelByID(ctx context.Context, channelID int) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("channel_id = ?", channelID).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	return &channel, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}
