package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

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

var wg sync.WaitGroup
var mykey string

var streamers = []Streamer{
	{Name: "Ice", ChannelId: "UCv9Edl_WbtbPeURPtFDo-uA"},
	{Name: "Mixhound", ChannelId: "UC_jxnWLGJ2eQK4en3UblKEw"},
	{Name: "Hyphonix", ChannelId: "UC4abN4ZiybnsAXTkTBX7now"},
	{Name: "Gary", ChannelId: "UCvxSwu13u1wWyROPlCH-MZg"},
	{Name: "Burger", ChannelId: "UC3MAdjjG3LMCG8CV-d7nEQA"},
	{Name: "Evan", ChannelId: "UCHYUiFsAJ-EDerAccSHIslw"},
	{Name: "Lolesports", ChannelId: "UCvqRdlKsE5Q8mf8YXbdIJLw"},
	{Name: "Chilledcow", ChannelId: "UCSJ4gkVC6NrvII8umztf0Ow"},
	{Name: "Cxnews", ChannelId: "UCStEQ9BjMLjHTHLNA6cY9vg"},
	{Name: "Code", ChannelId: "UCvjgXvBlbQiydffZU7m1_aw"},
	{Name: "Joe", ChannelId: "UCzQUP1qoWDoEbmsQxvdjxgQ"},
	{Name: "Nasa", ChannelId: "UCLA_DiR1FfKNvjuUpBHmylQ"},
	{Name: "CBS", ChannelId: "UC8p1vwvWtl6T73JiExfWs1g"},
	{Name: "Pepper", ChannelId: "UCdSr4xliU8yDyS1aGnCUMTA"},
}
var resp []Newlive

func getCatalog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	b, err := json.Marshal(streamers)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}
func (a ByViewers) Len() int      { return len(a) }
func (a ByViewers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByViewers) Less(i, j int) bool {
	return a[i].Viewers > a[j].Viewers
}

func sendStuff(w http.ResponseWriter, r *http.Request) {
	sort.Sort(ByViewers(resp))
	w.Header().Set("Content-type", "application/json")
	b, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func init() {
	ky := &mykey
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	*ky = os.Getenv("KEY")
	fmt.Println(mykey)
}

func main() {
	fmt.Println("Server Started...")
	go getter()
	go interval()
	r := mux.NewRouter()
	r.HandleFunc("/streamers/all", getCatalog).Methods("GET")
	r.HandleFunc("/streamers/live", sendStuff).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/build")))
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getter() {
	var results []Islive
	fmt.Println("getting....")
	for _, v := range streamers {
		url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&channelId=%v&eventType=live&type=video&key=%v", v.ChannelId, mykey)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		var streamer Islive
		json.Unmarshal(body, &streamer)
		if streamer.PageInfo.TotalResults == 0 {
			continue
		}
		streamer.Name = v.Name
		results = append(results, streamer)
	}
	var liveresults []Newlive
	for _, v := range results {
		id := v.Items[0].ID.VideoID
		resp, err := http.Get("https://www.googleapis.com/youtube/v3/videos?part=statistics%2C+snippet%2C+liveStreamingDetails&id=" + id + "&key=" + mykey)
		if err != nil {
			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var live Livestream
		json.Unmarshal(body, &live)
		name, err := strconv.Atoi(live.Items[0].LiveStreamingDetails.ConcurrentViewers)
		if err != nil {
			fmt.Println(err)
		}
		rz := Newlive{
			Name:        v.Name,
			ChannelID:   live.Items[0].Snippet.ChannelID,
			Title:       live.Items[0].Snippet.Title,
			Description: live.Items[0].Snippet.Description,
			Viewers:     name,
			Likes:       live.Items[0].Statistics.LikeCount,
			Dislikes:    live.Items[0].Statistics.DislikeCount,
			VideoID:     live.Items[0].ID,
		}
		liveresults = append(liveresults, rz)
	}
	resp = liveresults

	sort.Sort(ByViewers(resp))
}

func interval() {
	pollInterval := 10

	timerCh := time.Tick(time.Duration(pollInterval) * time.Minute)

	for range timerCh {
		getter()
	}
}
