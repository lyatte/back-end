// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Video struct {
	VideoID          string `json:"videoId"`
	VideoTitle       string `json:"videoTitle"`
	VideoDescription string `json:"videoDescription"`
	VideoCategory    string `json:"videoCategory"`
	VideoLike        int    `json:"videoLike"`
	VideoDislike     int    `json:"videoDislike"`
	VideoPrivacy     string `json:"videoPrivacy"`
	VideoPremium     bool   `json:"videoPremium"`
	VideoRestriction bool   `json:"videoRestriction"`
	VideoThumbnail   string `json:"videoThumbnail"`
	Video            string `json:"video"`
	ChannelID        int    `json:"channelId"`
	VideoViews       int    `json:"videoViews"`
	VideoRegion      string `json:"videoRegion"`
	Day              int    `json:"day"`
	Month            int    `json:"month"`
	Year             int    `json:"year"`
}

type NewVideo struct {
	VideoTitle       string `json:"videoTitle"`
	VideoDescription string `json:"videoDescription"`
	VideoCategory    string `json:"videoCategory"`
	VideoLike        int    `json:"videoLike"`
	VideoDislike     int    `json:"videoDislike"`
	VideoPrivacy     string `json:"videoPrivacy"`
	VideoPremium     bool   `json:"videoPremium"`
	VideoRestriction bool   `json:"videoRestriction"`
	VideoThumbnail   string `json:"videoThumbnail"`
	Video            string `json:"video"`
	ChannelID        int    `json:"channelId"`
	VideoViews       int    `json:"videoViews"`
	VideoRegion      string `json:"videoRegion"`
	Day              int    `json:"day"`
	Month            int    `json:"month"`
	Year             int    `json:"year"`
}
