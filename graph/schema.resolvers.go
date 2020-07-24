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
		ChannelName:      input.ChannelName,
		ChannelIcon:      input.ChannelIcon,
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

func (r *mutationResolver) AddVideoViews(ctx context.Context, videoID string) (bool, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("video_id = ?", videoID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("Video isn't valid")
	}

	video.VideoViews += 1

	_, updateError := r.DB.Model(&video).Where("video_id = ?", videoID).Update()

	if updateError != nil {
		return false, errors.New("Add Views failed")
	}

	return true, nil
}

func (r *mutationResolver) AddVideoLike(ctx context.Context, videoID string, channelID string) (bool, error) {
	var video model.Video
	var channel model.Channel

	err := r.DB.Model(&video).Where("video_id = ?", videoID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("Video isn't valid")
	}

	chErr := r.DB.Model(&channel).Where("channel_id = ?", channelID).Select()

	if chErr != nil {
		log.Println(chErr)
		return false, errors.New("Channel isn't valid")
	}

	video.VideoLike += 1

	channel.ChannelLikedVideo += videoID + ","

	_, updateErrorV := r.DB.Model(&video).Where("video_id = ?", videoID).Update()

	if updateErrorV != nil {
		return false, errors.New("Add like failed")
	}

	_, updateErrorC := r.DB.Model(&channel).Where("channel_id = ?", channelID).Update()

	if updateErrorC != nil {
		return false, errors.New("update channel failed")
	}

	return true, nil
}

func (r *mutationResolver) AddVideoDislike(ctx context.Context, videoID string, channelID string) (bool, error) {
	var video model.Video
	var channel model.Channel

	err := r.DB.Model(&video).Where("video_id = ?", videoID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("Video isn't valid")
	}

	chErr := r.DB.Model(&channel).Where("channel_id = ?", channelID).Select()

	if chErr != nil {
		log.Println(chErr)
		return false, errors.New("Channel isn't valid")
	}

	video.VideoDislike += 1

	channel.ChannelDislikedVideo += videoID + ","

	_, updateErrorV := r.DB.Model(&video).Where("video_id = ?", videoID).Update()

	if updateErrorV != nil {
		return false, errors.New("Add dislike failed")
	}

	_, updateErrorC := r.DB.Model(&channel).Where("channel_id = ?", channelID).Update()

	if updateErrorC != nil {
		return false, errors.New("update channel failed")
	}

	return true, nil
}

func (r *mutationResolver) CreateChannel(ctx context.Context, channelID string, input *model.NewChannel) (*model.Channel, error) {
	var ch model.Channel

	error := r.DB.Model(&ch).Where("channel_id = ?", channelID).Select()

	if error != nil {
		log.Println("ID still not registered!")
	} else {
		return nil, errors.New("ID already valid!")
	}

	channel := model.Channel{
		ChannelID:              input.ChannelID,
		ChannelName:            input.ChannelName,
		ChannelBackground:      input.ChannelBackground,
		ChannelIcon:            input.ChannelIcon,
		ChannelSubscribers:     input.ChannelSubscribers,
		ChannelDescription:     input.ChannelDescription,
		ChannelJoinDateDay:     input.ChannelJoinDateDay,
		ChannelJoinDateMonth:   input.ChannelJoinDateMonth,
		ChannelJoinDateYear:    input.ChannelJoinDateYear,
		ChannelLikedVideo:      input.ChannelLikedVideo,
		ChannelDislikedVideo:   input.ChannelDislikedVideo,
		ChannelLikedPost:       input.ChannelLikedPost,
		ChannelDislikedPost:    input.ChannelDislikedPost,
		ChannelLikedComment:    input.ChannelLikedComment,
		ChannelDislikedComment: input.ChannelDislikedComment,
		ChannelPremium:         input.ChannelPremium,
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

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input *model.NewPlaylist) (*model.Playlist, error) {

	playlist := model.Playlist{
		ChannelID:            input.ChannelID,
		PlaylistTitle:      input.PlaylistTitle,
		PlaylistDay:            input.PlaylistDay,
		PlaylistVisibility: input.PlaylistVisibility,
		PlaylistMonth:     input.PlaylistMonth,
		PlaylistYear:     input.PlaylistYear,
		PlaylistViews:     input.PlaylistViews,
		PlaylistVideos:   input.PlaylistVideos,
		PlaylistDesc:    input.PlaylistDesc,
	}

	_, err := r.DB.Model(&playlist).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert failed!")
	} else {
		log.Println("Insert success!")
		return &playlist, nil
	}
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, playlistID string, input *model.NewPlaylist) (*model.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, playlistID string) (bool, error) {
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

func (r *queryResolver) GetVideoByCategory(ctx context.Context, videoCategory string) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Where("video_category = ?", videoCategory).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	return videos, nil
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

func (r *queryResolver) GetChannelByID(ctx context.Context, channelID string) (*model.Channel, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("channel_id = ?", channelID).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	return &channel, nil
}

func (r *queryResolver) GetChannelPlaylist(ctx context.Context, channelID string) ([]*model.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
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
