package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"song_app/configs"
	"song_app/internal/models"
)

func (h *HTTPSongHandle) GetSongInfo(input *models.SongInput) (song *models.SongDetails, err error) {
	log.Println("GetSongInfo app started")
	conf, err := configs.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("can't get config data: %s", err)
	}

	params := url.Values{}
	params.Add("group", input.Group)
	params.Add("song", input.Song)

	url := conf.URL + params.Encode()

	ctx := context.Background()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error with NewRequestWithContext: %s", err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when executing the request to API: %s", err)
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("fail to read response: %s", err)
	}

	err = json.Unmarshal(content, &song)
	if err != nil {
		return nil, fmt.Errorf("corrupt json data: %s", err)
	}

	return song, err
}
