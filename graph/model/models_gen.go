// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Channel struct {
	ChannelID            string `json:"channel_id"`
	ChannelName          string `json:"channel_name"`
	ChannelBackground    string `json:"channel_background"`
	ChannelIcon          string `json:"channel_icon"`
	ChannelSubscribers   int    `json:"channel_subscribers"`
	ChannelDescription   string `json:"channel_description"`
	ChannelJoinDateDay   int    `json:"channel_join_date_day"`
	ChannelJoinDateMonth int    `json:"channel_join_date_month"`
	ChannelJoinDateYear  int    `json:"channel_join_date_year"`
}

type Video struct {
	VideoID          string `json:"video_id"`
	VideoTitle       string `json:"video_title"`
	VideoDescription string `json:"video_description"`
	VideoCategory    string `json:"video_category"`
	VideoLike        int    `json:"video_like"`
	VideoDislike     int    `json:"video_dislike"`
	VideoPrivacy     string `json:"video_privacy"`
	VideoPremium     bool   `json:"video_premium"`
	VideoRestriction bool   `json:"video_restriction"`
	VideoThumbnail   string `json:"video_thumbnail"`
	Video            string `json:"video"`
	VideoViews       int    `json:"video_views"`
	VideoRegion      string `json:"video_region"`
	Day              int    `json:"day"`
	Month            int    `json:"month"`
	Year             int    `json:"year"`
	ChannelID        int    `json:"channel_id"`
	ChannelName      string `json:"channel_name"`
	ChannelIcon      string `json:"channel_icon"`
}

type NewChannel struct {
	ChannelName          string `json:"channel_name"`
	ChannelBackground    string `json:"channel_background"`
	ChannelIcon          string `json:"channel_icon"`
	ChannelSubscribers   int    `json:"channel_subscribers"`
	ChannelDescription   string `json:"channel_description"`
	ChannelJoinDateDay   int    `json:"channel_join_date_day"`
	ChannelJoinDateMonth int    `json:"channel_join_date_month"`
	ChannelJoinDateYear  int    `json:"channel_join_date_year"`
}

type NewVideo struct {
	VideoTitle       string `json:"video_title"`
	VideoDescription string `json:"video_description"`
	VideoCategory    string `json:"video_category"`
	VideoLike        int    `json:"video_like"`
	VideoDislike     int    `json:"video_dislike"`
	VideoPrivacy     string `json:"video_privacy"`
	VideoPremium     bool   `json:"video_premium"`
	VideoRestriction bool   `json:"video_restriction"`
	VideoThumbnail   string `json:"video_thumbnail"`
	Video            string `json:"video"`
	VideoViews       int    `json:"video_views"`
	VideoRegion      string `json:"video_region"`
	Day              int    `json:"day"`
	Month            int    `json:"month"`
	Year             int    `json:"year"`
	ChannelID        int    `json:"channel_id"`
	ChannelName      string `json:"channel_name"`
	ChannelIcon      string `json:"channel_icon"`
}
