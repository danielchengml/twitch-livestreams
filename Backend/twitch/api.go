package twitchapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Streams struct {
	Total      int       `json:"_total"`
	LiveStream []*Stream `json:"streams"`
}

type Stream struct {
	Id           int      `json:"_id"`
	Game         string   `json:"game"`
	Viewers      int      `json:"viewers"`
	Video_height int      `json:"video_height"`
	Average_fps  float32  `json:"average_fps"`
	Delay        int      `json:"delay"`
	Created_at   string   `json:"created_at"`
	Is_playlist  bool     `json:"is_playlist"`
	Stream_type  string   `json:"stream_type"`
	Preview      *Preview `json:"preview"`
	Channel      *Channel `json:"channel"`
}

type Preview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}

type Channel struct {
	Mature                          bool    `json:"mature"`
	Partner                         bool    `json:"partner"`
	Status                          string  `json:"status"`
	Broadcaster_language            string  `json:"broadcaster_language"`
	Display_name                    string  `json:"display_name"`
	Game                            string  `json:"game"`
	Language                        string  `json:"language"`
	Id                              int     `json:"_id"`
	Name                            string  `json:"name"`
	Created_at                      string  `json:"created_at"`
	Updated_at                      string  `json:"updated_at"`
	Delay                           *string `json:"delay"`
	Logo                            string  `json:"logo"`
	Banner                          *string `json:"banner"`
	Video_banner                    string  `json:"video_banner"`
	Background                      *string `json:"background"`
	Profile_banner                  string  `json:"profile_banner"`
	Profile_banner_background_color string  `json:"profile_banner_background_color"`
	Url                             string  `json:"url"`
	Views                           int     `json:"views"`
	Followers                       int     `json:"followers"`
	Links                           *Links  `json:"_links"`
}

type Links struct {
	Self          string `json:"self"`
	Follows       string `json:"follows"`
	Commercial    string `json:"commercial"`
	Stream_key    string `json:"stream_key"`
	Chat          string `json:"chat"`
	Features      string `json:"features"`
	Subscriptions string `json:"subscriptions"`
	Editors       string `json:"editors"`
	Teams         string `json:"teams"`
	Videos        string `json:"videos"`
}

func getStreamlist(body []byte) (*Streams, error) {
	s := &Streams{}
	err := json.Unmarshal(body, s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func GetLivestreams(query string) (*Streams, error) {
	u, err := url.Parse("https://api.twitch.tv/kraken/search/streams")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("client_id", "hqtr6w669ykfhj1oawpvh6fgnizcwh")
	q.Set("limit", "100")
	q.Set("query", query)

	u.RawQuery = q.Encode()
	fmt.Printf("%v", u)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s, err := getStreamlist(body)

	return s, err
}
