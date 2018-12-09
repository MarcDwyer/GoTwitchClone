package types

import "time"

type Streamer struct {
	Name      string
	ChannelId string
}
type Islive struct {
	Kind       string `json:"kind"`
	Etag       string `json:"etag"`
	Name       string
	RegionCode string `json:"regionCode"`
	PageInfo   struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Etag string `json:"etag"`
		ID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
		Kind    string `json:"kind"`
		Snippet struct {
			ChannelID            string    `json:"channelId"`
			ChannelTitle         string    `json:"channelTitle"`
			Description          string    `json:"description"`
			LiveBroadcastContent string    `json:"liveBroadcastContent"`
			PublishedAt          time.Time `json:"publishedAt"`
			Thumbnails           struct {
				Default struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"default"`
				High struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"high"`
				Medium struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"medium"`
			} `json:"thumbnails"`
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
}

type Youtube struct {
	Kind       string `json:"kind"`
	Etag       string `json:"etag"`
	RegionCode string `json:"regionCode"`
	PageInfo   struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []interface{} `json:"items"`
}

type Livestream struct {
	Kind     string `json:"kind"`
	Etag     string `json:"etag"`
	Name     string
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
				Standard struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"standard"`
				Maxres struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"maxres"`
			} `json:"thumbnails"`
			ChannelTitle         string   `json:"channelTitle"`
			Tags                 []string `json:"tags"`
			CategoryID           string   `json:"categoryId"`
			LiveBroadcastContent string   `json:"liveBroadcastContent"`
			Localized            struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"localized"`
		} `json:"snippet"`
		Statistics struct {
			ViewCount     string `json:"viewCount"`
			LikeCount     string `json:"likeCount"`
			DislikeCount  string `json:"dislikeCount"`
			FavoriteCount string `json:"favoriteCount"`
		} `json:"statistics"`
		LiveStreamingDetails struct {
			ActualStartTime   time.Time `json:"actualStartTime"`
			ConcurrentViewers string    `json:"concurrentViewers"`
		} `json:"liveStreamingDetails"`
	} `json:"items"`
}
type Newlive struct {
	Name        string
	ChannelID   string
	Title       string
	Description string
	Viewers     int
	Likes       string
	Dislikes    string
	VideoID     string
}
type ByViewers []Newlive

func (a ByViewers) Len() int      { return len(a) }
func (a ByViewers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByViewers) Less(i, j int) bool {
	return a[i].Viewers > a[j].Viewers
}
