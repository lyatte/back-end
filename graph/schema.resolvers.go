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
	"math/rand"
	"sort"
	"strings"
	"time"
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

func (r *mutationResolver) AddChannelSubscribe(ctx context.Context, channelID string, chSubs string) (bool, error) {
	var channel model.Channel

	var dummy model.Channel

	err := r.DB.Model(&channel).Where("channel_id = ?", channelID).Select()

	if err != nil {
		log.Println("asd", err)
		return false, errors.New("asd Channel not valid!")
	}

	//playlist.PlaylistVideos += videoID + ","

	err1 := r.DB.Model(&dummy).Where("channel_id = ?", chSubs).Select()

	if err1 != nil {
		log.Println(err1)
		return false, errors.New("Channel not valid!")
	}

	channel.ChannelSubscribe += chSubs + ","

	dummy.ChannelSubscribers += 1

	_, updateError1 := r.DB.Model(&channel).Where("channel_id = ?", channelID).Update()

	if updateError1 != nil {
		return false, errors.New("Add subscribe failed")
	}

	_, updateError2 := r.DB.Model(&dummy).Where("channel_id = ?", chSubs).Update()

	if updateError2 != nil {
		return false, errors.New("Add subscriber failed")
	}

	return true, nil
}

func (r *mutationResolver) UnsubscribeChannel(ctx context.Context, channelID string, chSubs string) (bool, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("channel_id = ?", channelID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("No CHannel!")
	}

	var ch model.Channel

	err1 := r.DB.Model(&ch).Where("channel_id = ?", chSubs).Select()

	if err1 != nil {
		log.Println(err1)
		return false, errors.New("No Channel!")
	}

	temp := strings.Split(channel.ChannelSubscribe, ",")

	var temp2 string

	for index, _ := range temp {
		if temp[index] != chSubs{
			temp2 += temp[index] + ","
		}
		log.Println(temp[index])
	}

	log.Println(temp2)

	channel.ChannelSubscribe = temp2
	ch.ChannelSubscribers -= 1

	_, updateError1 := r.DB.Model(&channel).Where("channel_id = ?", channelID).Update()

	if updateError1 != nil {
		return false, errors.New("Change subscribe failed")
	}

	_, updateError2 := r.DB.Model(&ch).Where("channel_id = ?", chSubs).Update()

	if updateError2 != nil {
		return false, errors.New("Change subscriber failed")
	}

	return true, nil
}

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input *model.NewPlaylist) (*model.Playlist, error) {
	playlist := model.Playlist{
		ChannelID:          input.ChannelID,
		PlaylistTitle:      input.PlaylistTitle,
		PlaylistDay:        input.PlaylistDay,
		PlaylistVisibility: input.PlaylistVisibility,
		PlaylistMonth:      input.PlaylistMonth,
		PlaylistYear:       input.PlaylistYear,
		PlaylistViews:      input.PlaylistViews,
		PlaylistVideos:     input.PlaylistVideos,
		PlaylistDesc:       input.PlaylistDesc,
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
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Select()

	if err != nil {
		return false, errors.New("playlist not found!")
	}

	_, deleteErr := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Delete()

	if deleteErr != nil {
		return false, errors.New("Delete playlist failed")
	}

	return true, nil
}

func (r *mutationResolver) AddVideoToPlaylist(ctx context.Context, playlistID string, videoID string) (bool, error) {
	var playlist model.Playlist
	var video model.Video

	err := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("Playlist isn't valid")
	}

	vErr := r.DB.Model(&video).Where("video_id = ?", videoID).Select()

	if vErr != nil {
		log.Println(vErr)
		return false, errors.New("Video isn't valid")
	}

	playlist.PlaylistVideos += videoID + ","

	dt := time.Now()

	playlist.PlaylistDay = dt.Day()
	playlist.PlaylistMonth = int(dt.Month())
	playlist.PlaylistYear = dt.Year()

	_, updateErrorV := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Update()

	if updateErrorV != nil {
		return false, errors.New("Add video to playlist failed")
	}

	return true, nil
}

func (r *mutationResolver) AddPlaylistViews(ctx context.Context, playlistID string) (bool, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("Playlist isn't valid")
	}

	playlist.PlaylistViews += 1

	_, updateError := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Update()

	if updateError != nil {
		return false, errors.New("Add Views failed")
	}

	return true, nil
}

func (r *mutationResolver) DeleteVideoPlaylist(ctx context.Context, playlistID string, videoID string) (bool, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("Playlist isn't valid")
	}

	temp := strings.Split(playlist.PlaylistVideos, ",")

	var newVid string

	for index, _ := range temp {
		if temp[index] == "" {
			continue
		} else if temp[index] == videoID {
			continue
		} else {
			newVid += temp[index] + ","
		}
	}

	log.Println(newVid)

	playlist.PlaylistVideos = newVid

	_, updateError := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Update()

	if updateError != nil {
		return false, errors.New("Change failed")
	}

	return true, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	comment := model.Comment{
		ChannelID:  input.ChannelID,
		Like:       input.Like,
		Dislike:    input.Dislike,
		ReplyTo:    input.ReplyTo,
		Content:    input.Content,
		ReplyCount: input.ReplyCount,
		VideoID:    input.VideoID,
		PostID:     input.PostID,
		Day:        input.Day,
		Month:      input.Month,
		Year:       input.Year,
	}

	if input.ReplyTo != 0 {
		var comment model.Comment
		err2 := r.DB.Model(&comment).Where("comment_id = ?", input.ReplyTo).Select()

		if err2 != nil {
			log.Println(err2)
			return nil, errors.New("Comment id not found")
		}

		comment.ReplyCount += 1

		_, updateErr := r.DB.Model(&comment).Where("comment_id = ?", input.ReplyTo).Update()

		if updateErr != nil {
			return nil, errors.New("Add reply count failed")
		}
	}

	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert failed!")
	} else {
		log.Println("Insert success!")
		return &comment, nil
	}
}

func (r *mutationResolver) UpdateCommentDl(ctx context.Context, commentID string, channelID string, flag int) (bool, error) {
	var comment model.Comment

	err := r.DB.Model(&comment).Where("comment_id = ?", commentID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("Comment isn't valid")
	}

	var channel model.Channel

	err2 := r.DB.Model(&channel).Where("channel_id = ?", channelID).Select()

	if err2 != nil {
		log.Println(err2)
		return false, errors.New("Comment isn't valid")
	}

	if flag == 1 {
		temp := strings.Split(channel.ChannelLikedComment, ",")

		var newArr string

		var flags int = 0

		for index, _ := range temp {
			if temp[index] == commentID {
				comment.Like -= 1
				flags += 1
			} else {
				newArr += temp[index] + ","
			}
		}

		if flags == 0 {
			comment.Like += 1

			channel.ChannelLikedComment += commentID + ","
		} else {
			channel.ChannelLikedComment = newArr
		}

	} else {
		temp := strings.Split(channel.ChannelDislikedComment, ",")

		var newArr string

		var flags int = 0

		for index, _ := range temp {
			if temp[index] == commentID {
				comment.Dislike -= 1
				flags += 1
			} else {
				newArr += temp[index] + ","
			}
		}

		if flags == 0 {
			comment.Dislike += 1

			channel.ChannelDislikedComment += commentID + ","
		} else {
			channel.ChannelDislikedComment = newArr
		}
	}

	_, updateError := r.DB.Model(&comment).Where("comment_id = ?", commentID).Update()

	if updateError != nil {
		return false, errors.New("Add Like/Dislike failed")
	}

	_, updateError2 := r.DB.Model(&channel).Where("channel_id = ?", channelID).Update()

	if updateError2 != nil {
		return false, errors.New("Add Like/Dislike failed")
	}

	return true, nil
}

func (r *mutationResolver) UpdateCommentRc(ctx context.Context, commentID string) (bool, error) {
	var comment model.Comment

	err := r.DB.Model(&comment).Where("comment_id = ?", commentID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("Comment isn't valid")
	}

	comment.ReplyCount += 1

	_, updateError := r.DB.Model(&comment).Where("comment_id = ?", commentID).Update()

	if updateError != nil {
		return false, errors.New("Add Reply Count failed")
	}

	return true, nil
}

func (r *mutationResolver) CreateMembership(ctx context.Context, input *model.NewMembership) (*model.Membership, error) {
	membership := model.Membership{
		MembershipID:    input.MembershipID,
		MembershipName:  input.MembershipName,
		MembershipPrice: input.MembershipPrice,
	}

	_, err := r.DB.Model(&membership).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert failed!")
	} else {
		log.Println("Insert success!")
		return &membership, nil
	}
}

func (r *mutationResolver) UpdateAccountPremium(ctx context.Context, channelID string, premiumID string, day int, month int, year int) (bool, error) {
	var channel model.Channel

	err := r.DB.Model(&channel).Where("channel_id = ?", channelID).Select()

	if err != nil {
		log.Println(err)
		return false, errors.New("Channel isn't valid")
	}

	var membership model.Membership

	err1 := r.DB.Model(&membership).Where("membership_id = ?", premiumID).Select()

	if err1 != nil {
		log.Println(err)
		return false, errors.New("Membership isn't valid")
	}

	channel.ChannelPremium = membership.MembershipID
	channel.PremiumDay = day
	channel.PremiumMonth = month
	channel.PremiumYear = year

	_, updateError := r.DB.Model(&channel).Where("channel_id = ?", channelID).Update()

	if updateError != nil {
		return false, errors.New("Premiumship failed")
	}

	return true, nil
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

func (r *queryResolver) GetVideosComment(ctx context.Context, videoID string, flag string) ([]*model.Comment, error) {
	var comment []*model.Comment

	err := r.DB.Model(&comment).Where("video_id = ?", videoID).Order("comment_id").Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}


	if flag == "1"{
		sort.Slice(comment[:], func(i, j int) bool {
			return comment[i].Like > comment[j].Like
		})
	}else{
		var temp []*model.Comment

		for i:= len(comment)-1; i>=0; i--{
			temp = append(temp, comment[i])
		}

		return temp, nil
	}

	return comment, nil
}

func (r *queryResolver) GetCommentReply(ctx context.Context, commentID string) ([]*model.Comment, error) {
	var comment []*model.Comment

	err := r.DB.Model(&comment).Where("reply_to = ?", commentID).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	return comment, nil
}

func (r *queryResolver) GetVideoByLocation(ctx context.Context, location string) ([]*model.Video, error) {
	var video []*model.Video
	var video2 []*model.Video

	//video, err ;= r.DB.Query("SELECT * FROM videos WHERE ")

	err := r.DB.Model(&video).Where("video_region != ?", location).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	err2 := r.DB.Model(&video2).Where("video_region = ?", location).Select()

	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("Query failed")
	}

	log.Println(video, video2)

	for index, _ := range video {
		video2 = append(video2, video[index])
	}

	return video2, nil
}

func (r *queryResolver) GetVideoByRestriction(ctx context.Context, restriction string) ([]*model.Video, error) {
	var video []*model.Video
	var video2 []*model.Video

	err := r.DB.Model(&video).Where("video_restriction = ?", "No").Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	if restriction == "No" {
		err := r.DB.Model(&video2).Where("video_restriction = ?", "Yes").Select()

		if err != nil {
			log.Println(err)
			return nil, errors.New("Query failed")
		}

		for index, _ := range video2 {
			video = append(video, video2[index])
		}

		return video, nil
	}

	log.Println(video, restriction)

	return video, nil
}

func (r *queryResolver) GetVideoHomePage(ctx context.Context, restriction string, location string, premiumID string) ([]*model.Video, error) {
	var vid []*model.Video
	var vid2 []*model.Video

	var video_res []*model.Video

	err_pr := r.DB.Model(&vid).Where("video_premium = ?", "false").Select()

	if err_pr != nil {
		log.Println(err_pr)
		return nil, errors.New("asd Query failed")
	}

	for index, _ := range vid {
		log.Println(vid[index], "Asd", len(vid))
	}

	if premiumID == "1" || premiumID == "2" {
		log.Println("masuk sini")
		err_pr2 := r.DB.Model(&vid2).Where("video_premium = ?", "true").Select()

		if err_pr2 != nil {
			log.Println(err_pr2)
			return nil, errors.New("Query failed")
		}

		for index, _ := range vid2 {
			vid = append(vid, vid2[index])
		}
	}

	log.Println("here2", len(vid))

	for index, _ := range vid {
		if vid[index].VideoRestriction == "No" {
			video_res = append(video_res, vid[index])
		}
	}

	log.Println("hereee", len(video_res))

	if restriction == "No" {
		log.Println("masuk sini desu")
		for index, _ := range vid {
			if vid[index].VideoRestriction == "Yes" {
				video_res = append(video_res, vid[index])
			}
		}

	}

	log.Println("here", len(video_res))

	var video2 []*model.Video

	for index, _ := range video_res {
		if video_res[index].VideoRegion == location {
			video2 = append(video2, video_res[index])
		}

	}

	var temp int

	var shuffled_vids []*model.Video

	var i int

	i = len(video2) - 1

	for {
		if i >= 0 {
			rand.Seed(time.Now().Unix())

			//log.Println(len(video2))

			if (len(video2) - 1) != 0 {
				temp = rand.Intn(len(video2) - 1)
				shuffled_vids = append(shuffled_vids, video2[temp])
			} else {
				shuffled_vids = append(shuffled_vids, video2[0])
			}

			//log.Println(temp)
			log.Println(i, len(video2), temp)

			video2[temp] = video2[len(video2)-1]
			video2[len(video2)-1] = nil
			video2 = video2[:len(video2)-1]

			i -= 1
			continue
		}
		break
	}

	var video3 []*model.Video

	for index, _ := range video_res {
		if video_res[index].VideoRegion != location {
			video3 = append(video3, video_res[index])
		}
	}

	i = len(video3) - 1

	var temp2 []*model.Video

	for {
		if i >= 0 {
			rand.Seed(time.Now().Unix())

			if (len(video3) - 1) != 0 {
				temp = rand.Intn(len(video3) - 1)
				temp2 = append(temp2, video3[temp])
			} else {
				temp2 = append(temp2, video3[0])
			}

			video3[temp] = video3[len(video3)-1]
			video3[len(video3)-1] = nil
			video3 = video3[:len(video3)-1]

			i -= 1
			continue
		}
		break
	}

	shuffled_vids = append(shuffled_vids, temp2...)

	return shuffled_vids, nil
}

func (r *queryResolver) GetVideoOrderedByViews(ctx context.Context) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Order("video_views").Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	date := time.Now()

	day := date.Day()
	month := int(date.Month())
	year := date.Year()

	month *= 30
	year *= 365

	var final_array []*model.Video

	for index, _ := range video {
		if (day+month+year)-(video[index].Day+video[index].Month*30+video[index].Year*365) < 7 {
			final_array = append(final_array, video[index])
		}
	}

	sort.Slice(final_array[:], func(i, j int) bool {
		return final_array[i].VideoViews > final_array[j].VideoViews
	})

	return final_array, nil
}

func (r *queryResolver) GetRelatedVideo(ctx context.Context, restriction string, premiumID string, location string, category string) ([]*model.Video, error) {
	var vid []*model.Video
	var vid2 []*model.Video

	var video_res []*model.Video

	err_pr := r.DB.Model(&vid).Where("video_premium = ?", "false").Select()

	if err_pr != nil {
		log.Println(err_pr)
		return nil, errors.New("asd Query failed")
	}

	if premiumID == "1" || premiumID == "2" {
		err_pr2 := r.DB.Model(&vid2).Where("video_premium = ?", "true").Select()

		if err_pr2 != nil {
			log.Println(err_pr2)
			return nil, errors.New("Query failed")
		}

		for index, _ := range vid2 {
			vid = append(vid, vid2[index])
		}
	}

	for index, _ := range vid {
		if vid[index].VideoRestriction == "No" {
			video_res = append(video_res, vid[index])
		}
	}
	if restriction == "No" {
		for index, _ := range vid {
			if vid[index].VideoRestriction == "Yes" {
				video_res = append(video_res, vid[index])
			}
		}
	}

	var video []*model.Video

	for index, _ := range video_res {
		if video_res[index].VideoRegion == location {
			video = append(video, video_res[index])
		}
	}

	var video2 []*model.Video

	for index, _ := range video_res {
		if video_res[index].VideoCategory == category {
			video2 = append(video2, video_res[index])
		}
	}

	var final_vids []*model.Video

	final_vids = append(final_vids, video...)

	final_vids = append(final_vids, video2...)

	return final_vids, nil
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
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Where("channel_id = ?", channelID).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	var p_vis []*model.Playlist

	for index, _ := range playlists {
		if playlists[index].PlaylistVisibility == "Public" {
			p_vis = append(p_vis, playlists[index])
		}
	}

	return p_vis, nil
}

func (r *queryResolver) GetChannelVideo(ctx context.Context, channelID string, flag string) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Where("channel_id = ?", channelID).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	var vid_prem []*model.Video

	for index, _ := range video {
		if video[index].VideoPremium == "false" {
			vid_prem = append(vid_prem, video[index])
		}
	}

	if flag == "1" || flag == "2" {
		for index, _ := range video {
			if video[index].VideoPremium == "true" {
				vid_prem = append(vid_prem, video[index])
			}
		}
	}

	return vid_prem, nil
}

func (r *queryResolver) GetChannelVideoMostPopular(ctx context.Context, channelID string, flag string) ([]*model.Video, error) {
	var video []*model.Video

	err1 := r.DB.Model(&video).Where("channel_id = ?", channelID).Select()

	if err1 != nil {
		log.Println(err1)
		return nil, errors.New("Query failed")
	}

	var video_prem []*model.Video

	for index, _ := range video {
		if video[index].VideoPremium == "false" {
			video_prem = append(video_prem, video[index])
		}
	}

	if flag == "1" || flag == "2" {
		for index, _ := range video {
			if video[index].VideoPremium == "true" {
				video_prem = append(video_prem, video[index])
			}
		}
	}

	sort.Slice(video_prem[:], func(i, j int) bool {
		return video_prem[i].VideoViews > video_prem[j].VideoViews
	})

	return video_prem, nil
}

func (r *queryResolver) GetChannelVideoDate(ctx context.Context, channelID string, prem string, flag string) ([]*model.Video, error) {
	var video []*model.Video

	err1 := r.DB.Model(&video).Where("channel_id = ?", channelID).Select()

	if err1 != nil {
		log.Println(err1)
		return nil, errors.New("Query failed")
	}

	var video_prem []*model.Video

	for index, _ := range video {
		if video[index].VideoPremium == "false" {
			video_prem = append(video_prem, video[index])
		}
	}

	if prem == "1" || prem == "2" {
		for index, _ := range video {
			if video[index].VideoPremium == "true" {
				video_prem = append(video_prem, video[index])
			}
		}
	}

	if flag == "1" { //newest
		sort.Slice(video_prem[:], func(i, j int) bool {
			return video_prem[i].Day+video_prem[i].Month*30+video_prem[i].Year*365 > video_prem[j].Day+video_prem[j].Month*30+video_prem[j].Year*365
		})
	} else { //oldest
		sort.Slice(video_prem[:], func(i, j int) bool {
			return video_prem[i].Day+video_prem[i].Month*30+video_prem[i].Year*365 < video_prem[j].Day+video_prem[j].Month*30+video_prem[j].Year*365
		})
	}

	return video_prem, nil
}

func (r *queryResolver) GetPlaylistByID(ctx context.Context, playlistID string) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("playlist_id = ?", playlistID).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	return &playlist, nil
}

func (r *queryResolver) GetSubscribeVideos(ctx context.Context, channelID []string, flag string) ([]*model.Video, error) {
	var channel model.Channel

	var final_array []*model.Video

	for index, _ := range channelID {
		if channelID[index] == "" {
			continue
		}
		err := r.DB.Model(&channel).Where("channel_id = ?", channelID[index]).Select()

		if err != nil {
			log.Println(err)
			return nil, errors.New("Query failed")
		}

		var video []*model.Video

		err1 := r.DB.Model(&video).Where("channel_id = ?", channelID[index]).Select()

		if err1 != nil {
			log.Println(err1)
			return nil, errors.New("Query failed")
		}

		date := time.Now()

		day := date.Day()
		month := int(date.Month())
		year := date.Year()

		if flag == "1" {
			for index, _ := range video {
				if video[index].Day == day && video[index].Month*30 == month*30 && video[index].Year*365 == year*365 {
					final_array = append(final_array, video[index])
				}
			}
		} else if flag == "2" {

			month *= 30
			year *= 365

			for index, _ := range video {
				if (day+month+year)-(video[index].Day+video[index].Month*30+video[index].Year*365) < 7 {
					log.Println((day + month + year) - (date.Day() + int(date.Month())*30 + date.Year()*365))
					final_array = append(final_array, video[index])
				}
			}

			sort.Slice(final_array[:], func(i, j int) bool {
				return final_array[i].Day+final_array[i].Month*30+final_array[i].Year*365 > final_array[j].Day+final_array[j].Month*30+final_array[j].Year*365
			})
		} else {
			month *= 30
			year *= 365

			for index, _ := range video {
				if (day+month+year)-(video[index].Day+video[index].Month*30+video[index].Year*365) < 30 {
					log.Println((day + month + year) - (date.Day() + int(date.Month())*30 + date.Year()*365))
					final_array = append(final_array, video[index])
				}
			}

			sort.Slice(final_array[:], func(i, j int) bool {
				return final_array[i].Day+final_array[i].Month*30+final_array[i].Year*365 > final_array[j].Day+final_array[j].Month*30+final_array[j].Year*365
			})
		}
	}

	return final_array, nil
}

func (r *queryResolver) GetMemberships(ctx context.Context) ([]*model.Membership, error) {
	var memberships []*model.Membership

	err := r.DB.Model(&memberships).Order("membership_id").Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	return memberships, nil
}

func (r *queryResolver) GetVideoCategoryAllTimePopular(ctx context.Context, restriction string, premium string, category string) ([]*model.Video, error) {
	var video []*model.Video

	err1 := r.DB.Model(&video).Where("video_category = ?", category).Select()

	if err1 != nil {
		log.Println(err1)
		return nil, errors.New("Query failed")
	}

	var video_res []*model.Video

	for index, _ := range video {
		if video[index].VideoRestriction == "No" {
			video_res = append(video_res, video[index])
		}
	}

	if restriction == "No" {
		for index, _ := range video {
			if video[index].VideoRestriction == "Yes" {
				video_res = append(video_res, video[index])
			}
		}
	}

	var video_prem []*model.Video

	for index, _ := range video_res {
		if video_res[index].VideoPremium == "false" {
			video_prem = append(video_prem, video_res[index])
		}
	}

	if premium == "1" || premium == "2" {
		for index, _ := range video_res {
			if video_res[index].VideoPremium == "true" {
				video_prem = append(video_prem, video_res[index])
			}
		}
	}

	sort.Slice(video_prem[:], func(i, j int) bool {
		return video_prem[i].VideoViews > video_prem[j].VideoViews
	})

	return video_prem, nil
}

func (r *queryResolver) GetVideoCategoryWeekPopular(ctx context.Context, restriction string, premium string, category string) ([]*model.Video, error) {
	var final_array []*model.Video

	var video []*model.Video

	err1 := r.DB.Model(&video).Where("video_category = ?", category).Select()

	if err1 != nil {
		log.Println(err1)
		return nil, errors.New("Query failed")
	}

	var video_res []*model.Video

	for index, _ := range video {
		if video[index].VideoRestriction == "No" {
			video_res = append(video_res, video[index])
		}
	}

	if restriction == "No" {
		for index, _ := range video {
			if video[index].VideoRestriction == "Yes" {
				video_res = append(video_res, video[index])
			}
		}
	}

	var video_prem []*model.Video

	for index, _ := range video_res {
		if video_res[index].VideoPremium == "false" {
			video_prem = append(video_prem, video_res[index])
		}
	}

	if premium == "1" || premium == "2" {
		for index, _ := range video_res {
			if video_res[index].VideoPremium == "true" {
				video_prem = append(video_prem, video_res[index])
			}
		}
	}

	date := time.Now()

	day := date.Day()
	month := int(date.Month())
	year := date.Year()

	month *= 30
	year *= 365

	for index, _ := range video_prem {
		if (day+month+year)-(video_prem[index].Day+video_prem[index].Month*30+video_prem[index].Year*365) < 7 {
			log.Println((day + month + year) - (date.Day() + int(date.Month())*30 + date.Year()*365))
			final_array = append(final_array, video_prem[index])
		}
	}

	sort.Slice(final_array[:], func(i, j int) bool {
		return final_array[i].VideoViews > final_array[j].VideoViews
	})

	return final_array, nil
}

func (r *queryResolver) GetVideoCategoryMonthPopular(ctx context.Context, restriction string, premium string, category string) ([]*model.Video, error) {
	var final_array []*model.Video

	var video []*model.Video

	err1 := r.DB.Model(&video).Where("video_category = ?", category).Select()

	if err1 != nil {
		log.Println(err1)
		return nil, errors.New("Query failed")
	}

	var video_res []*model.Video

	for index, _ := range video {
		if video[index].VideoRestriction == "No" {
			video_res = append(video_res, video[index])
		}
	}

	if restriction == "No" {
		for index, _ := range video {
			if video[index].VideoRestriction == "Yes" {
				video_res = append(video_res, video[index])
			}
		}
	}

	var video_prem []*model.Video

	for index, _ := range video_res {
		if video_res[index].VideoPremium == "false" {
			video_prem = append(video_prem, video_res[index])
		}
	}

	if premium == "1" || premium == "2" {
		for index, _ := range video_res {
			if video_res[index].VideoPremium == "true" {
				video_prem = append(video_prem, video_res[index])
			}
		}
	}

	date := time.Now()

	day := date.Day()
	month := int(date.Month())
	year := date.Year()

	month *= 30
	year *= 365

	for index, _ := range video_prem {
		if (day+month+year)-(video_prem[index].Day+video_prem[index].Month*30+video_prem[index].Year*365) < 30 {
			log.Println((day + month + year) - (date.Day() + int(date.Month())*30 + date.Year()*365))
			final_array = append(final_array, video_prem[index])
		}
	}

	sort.Slice(final_array[:], func(i, j int) bool {
		return final_array[i].VideoViews > final_array[j].VideoViews
	})

	return final_array, nil
}

func (r *queryResolver) GetVideoCategoryRecently(ctx context.Context, restriction string, premium string, category string) ([]*model.Video, error) {
	var video []*model.Video

	err1 := r.DB.Model(&video).Where("video_category = ?", category).Select()

	if err1 != nil {
		log.Println(err1)
		return nil, errors.New("Query failed")
	}

	var video_res []*model.Video

	for index, _ := range video {
		if video[index].VideoRestriction == "No" {
			video_res = append(video_res, video[index])
		}
	}

	if restriction == "No" {
		for index, _ := range video {
			if video[index].VideoRestriction == "Yes" {
				video_res = append(video_res, video[index])
			}
		}
	}

	var video_prem []*model.Video

	for index, _ := range video_res {
		if video_res[index].VideoPremium == "false" {
			video_prem = append(video_prem, video_res[index])
		}
	}

	if premium == "1" || premium == "2" {
		for index, _ := range video_res {
			if video_res[index].VideoPremium == "true" {
				video_prem = append(video_prem, video_res[index])
			}
		}
	}

	sort.Slice(video_prem[:], func(i, j int) bool {
		return video_prem[i].Day+video_prem[i].Month*30+video_prem[i].Year*365 > video_prem[j].Day+video_prem[j].Month*30+video_prem[j].Year*365
	})

	return video_prem, nil
}

func (r *queryResolver) GetSearchVideo(ctx context.Context, keyword string, uploadDate string, premium string) ([]*model.Video, error) {
	var video []*model.Video

	key := "%" + keyword + "%"

	err := r.DB.Model(&video).Where("LOWER(video_title) LIKE LOWER(?)", key).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	var video_res []*model.Video

	for index, _ := range video {
		if video[index].VideoPremium == "false" {
			video_res = append(video_res, video[index])
		}
	}

	if premium == "1" || premium == "2" {
		for index, _ := range video {
			if video[index].VideoPremium == "true" {
				video_res = append(video_res, video[index])
			}
		}
	}

	date := time.Now()

	day := date.Day()
	month := int(date.Month())
	year := date.Year()

	var final_array []*model.Video

	if uploadDate == "1" {
		month *= 30
		year *= 365

		for index, _ := range video_res {
			if (day+month+year)-(video_res[index].Day+video_res[index].Month*30+video_res[index].Year*365) < 7 {
				final_array = append(final_array, video_res[index])
			}
		}

		sort.Slice(final_array[:], func(i, j int) bool {
			return final_array[i].Day+final_array[i].Month*30+final_array[i].Year*365 > final_array[j].Day+final_array[j].Month*30+final_array[j].Year*365
		})
	} else if uploadDate == "2" {
		month *= 30
		year *= 365

		for index, _ := range video_res {
			if (day+month+year)-(video_res[index].Day+video_res[index].Month*30+video_res[index].Year*365) < 30 {
				final_array = append(final_array, video_res[index])
			}
		}

		sort.Slice(final_array[:], func(i, j int) bool {
			return final_array[i].Day+final_array[i].Month*30+final_array[i].Year*365 > final_array[j].Day+final_array[j].Month*30+final_array[j].Year*365
		})

	} else if uploadDate == "3"{
		month *= 30
		year *= 365

		for index, _ := range video_res {
			if (day+month+year)-(video_res[index].Day+video_res[index].Month*30+video_res[index].Year*365) < 365 {
				final_array = append(final_array, video_res[index])
			}
		}

		sort.Slice(final_array[:], func(i, j int) bool {
			return final_array[i].Day+final_array[i].Month*30+final_array[i].Year*365 > final_array[j].Day+final_array[j].Month*30+final_array[j].Year*365
		})
	} else {
		final_array = append(final_array, video_res...)

		sort.Slice(final_array[:], func(i, j int) bool {
			return final_array[i].Day+final_array[i].Month*30+final_array[i].Year*365 > final_array[j].Day+final_array[j].Month*30+final_array[j].Year*365
		})
	}

	return final_array, nil
}

func (r *queryResolver) GetSearchPlaylist(ctx context.Context, keyword string, uploadDate string) ([]*model.Playlist, error) {
	var playlist []*model.Playlist

	key := "%" + keyword + "%"

	err := r.DB.Model(&playlist).Where("LOWER(playlist_title) LIKE LOWER(?)", key).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	date := time.Now()

	day := date.Day()
	month := int(date.Month())
	year := date.Year()

	var final_array []*model.Playlist

	if uploadDate == "1" {
		month *= 30
		year *= 365

		for index, _ := range playlist {
			if (day+month+year)-(playlist[index].PlaylistDay+playlist[index].PlaylistMonth*30+playlist[index].PlaylistYear*365) < 7 {
				final_array = append(final_array, playlist[index])
			}
		}

		sort.Slice(final_array[:], func(i, j int) bool {
			return final_array[i].PlaylistDay+final_array[i].PlaylistMonth*30+final_array[i].PlaylistYear*365 > final_array[j].PlaylistDay+final_array[j].PlaylistMonth*30+final_array[j].PlaylistYear*365
		})
	} else if uploadDate == "2" {
		month *= 30
		year *= 365

		for index, _ := range playlist {
			if (day+month+year)-(playlist[index].PlaylistDay+playlist[index].PlaylistMonth*30+playlist[index].PlaylistYear*365) < 30 {
				final_array = append(final_array, playlist[index])
			}
		}

		sort.Slice(final_array[:], func(i, j int) bool {
			return final_array[i].PlaylistDay+final_array[i].PlaylistMonth*30+final_array[i].PlaylistYear*365 > final_array[j].PlaylistDay+final_array[j].PlaylistMonth*30+final_array[j].PlaylistYear*365
		})

	} else {
		month *= 30
		year *= 365

		for index, _ := range playlist {
			if (day+month+year)-(playlist[index].PlaylistDay+playlist[index].PlaylistMonth*30+playlist[index].PlaylistYear*365) < 365 {
				final_array = append(final_array, playlist[index])
			}
		}

		sort.Slice(final_array[:], func(i, j int) bool {
			return final_array[i].PlaylistDay+final_array[i].PlaylistMonth*30+final_array[i].PlaylistYear*365 > final_array[j].PlaylistDay+final_array[j].PlaylistMonth*30+final_array[j].PlaylistYear*365
		})
	}

	return final_array, nil
}

func (r *queryResolver) GetSearchChannel(ctx context.Context, keyword string, uploadDate string) ([]*model.Channel, error) {
	var channel []*model.Channel

	key := "%" + keyword + "%"

	err := r.DB.Model(&channel).Where("LOWER(channel_name) LIKE LOWER(?)", key).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Query failed")
	}

	return channel, nil
}

func (r *queryResolver) GetChannelVideoRecent(ctx context.Context, channelID string, flag string) ([]*model.Video, error) {
	var video []*model.Video

	err1 := r.DB.Model(&video).Where("channel_id = ?", channelID).Select()

	if err1 != nil {
		log.Println(err1)
		return nil, errors.New("Query failed")
	}

	var video_prem []*model.Video

	for index, _ := range video {
		if video[index].VideoPremium == "false" {
			video_prem = append(video_prem, video[index])
		}
	}

	if flag == "1" || flag == "2" {
		for index, _ := range video {
			if video[index].VideoPremium == "true" {
				video_prem = append(video_prem, video[index])
			}
		}
	}

	sort.Slice(video_prem[:], func(i, j int) bool {
		return video_prem[i].Day+video_prem[i].Month*30+video_prem[i].Year*365 > video_prem[j].Day+video_prem[j].Month*30+video_prem[j].Year*365
	})

	return video_prem, nil
}

func (r *queryResolver) GetChannelRandomVideo(ctx context.Context, channelID string, flag string) ([]*model.Video, error) {
	var vid []*model.Video

	var video_prem []*model.Video

	err := r.DB.Model(&vid).Where("channel_id = ?", channelID).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("asd Query failed")
	}

	for index, _ := range vid {
		if vid[index].VideoPremium == "false" {
			video_prem = append(video_prem, vid[index])
		}
	}

	if flag == "1" || flag == "2" {
		for index, _ := range vid {
			if vid[index].VideoPremium == "true" {
				video_prem = append(video_prem, vid[index])
			}
		}
	}

	var temp int

	var shuffled_vids []*model.Video

	var i int

	i = len(video_prem) - 1

	for {
		if i >= 0 {
			rand.Seed(time.Now().Unix())

			//log.Println(len(video2))

			if (len(video_prem) - 1) != 0 {
				temp = rand.Intn(len(video_prem) - 1)
				shuffled_vids = append(shuffled_vids, video_prem[temp])
			} else {
				shuffled_vids = append(shuffled_vids, video_prem[0])
			}

			video_prem[temp] = video_prem[len(video_prem)-1]
			video_prem[len(video_prem)-1] = nil
			video_prem = video_prem[:len(video_prem)-1]

			i -= 1
			continue
		}
		break
	}

	return shuffled_vids, nil
}

func (r *queryResolver) GetChannelRandomPlaylist(ctx context.Context, channelID string) ([]*model.Playlist, error) {
	var playlist []*model.Playlist

	err := r.DB.Model(&playlist).Where("channel_id = ?", channelID).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("asd Query failed")
	}

	var playlist2 []*model.Playlist

	for index, _ := range playlist {
		if playlist[index].PlaylistVisibility == "Public" {
			playlist2 = append(playlist2, playlist[index])
		}
	}

	var temp int

	var shuffled_p []*model.Playlist

	var i int

	i = len(playlist2) - 1

	for {
		if i >= 0 {
			rand.Seed(time.Now().Unix())

			//log.Println(len(video2))

			if (len(playlist2) - 1) != 0 {
				temp = rand.Intn(len(playlist2) - 1)
				shuffled_p = append(shuffled_p, playlist2[temp])
			} else {
				shuffled_p = append(shuffled_p, playlist2[0])
			}

			playlist2[temp] = playlist2[len(playlist2)-1]
			playlist2[len(playlist2)-1] = nil
			playlist2 = playlist2[:len(playlist2)-1]

			i -= 1
			continue
		}
		break
	}

	return shuffled_p, nil
}

func (r *queryResolver) GetPlaylistVideo(ctx context.Context, videos string, flag string) ([]*model.Video, error) {
	var vid []*model.Video

	err := r.DB.Model(&vid).Order("channel_id").Select()

	if err != nil {
		return nil, errors.New("asd Query failed")
	}

	temp := strings.Split(videos, ",")

	var fin_vids []*model.Video

	for index, _ := range temp {
		for i, _ := range vid {
			if temp[index] == "" {
				continue
			}
			log.Println(temp[index])

			if temp[index] == vid[i].VideoID {
				fin_vids = append(fin_vids, vid[i])
			}
		}
	}

	var t []*model.Video

	if flag == "1" { //newest
		return fin_vids, nil
	} else { //oldest
		var temp = len(fin_vids) - 1

		for i := temp; i >= 0; i-- {
			t = append(t, fin_vids[i])
		}

		return t, nil
	}
}

func (r *queryResolver) GetPlaylistVideoDp(ctx context.Context, videos string, flag string) ([]*model.Video, error) {
	var vid []*model.Video

	err := r.DB.Model(&vid).Order("channel_id").Select()

	if err != nil {
		return nil, errors.New("asd Query failed")
	}

	temp := strings.Split(videos, ",")

	var fin_vids []*model.Video

	for index, _ := range temp {
		for i, _ := range vid {
			if temp[index] == "" {
				continue
			}
			log.Println(temp[index])

			if temp[index] == vid[i].VideoID {
				fin_vids = append(fin_vids, vid[i])
			}
		}
	}

	if flag == "1" { //newest
		sort.Slice(fin_vids[:], func(i, j int) bool {
			return fin_vids[i].Day+fin_vids[i].Month*30+fin_vids[i].Year*365 > fin_vids[j].Day+fin_vids[j].Month*30+fin_vids[j].Year*365
		})
	} else { //oldest
		sort.Slice(fin_vids[:], func(i, j int) bool {
			return fin_vids[i].Day+fin_vids[i].Month*30+fin_vids[i].Year*365 < fin_vids[j].Day+fin_vids[j].Month*30+fin_vids[j].Year*365
		})
	}

	return fin_vids, nil
}

func (r *queryResolver) GetPlaylistVideoPopularity(ctx context.Context, videos string) ([]*model.Video, error) {
	var vid []*model.Video

	err := r.DB.Model(&vid).Order("channel_id").Select()

	if err != nil {
		return nil, errors.New("asd Query failed")
	}

	temp := strings.Split(videos, ",")

	var fin_vids []*model.Video

	for index, _ := range temp {
		for i, _ := range vid {
			if temp[index] == "" {
				continue
			}
			log.Println(temp[index])

			if temp[index] == vid[i].VideoID {
				fin_vids = append(fin_vids, vid[i])
			}
		}
	}

	sort.Slice(fin_vids[:], func(i, j int) bool {
		return fin_vids[i].VideoViews > fin_vids[j].VideoViews
	})

	return fin_vids, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
