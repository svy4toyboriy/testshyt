package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	apiKey    = "KEY"
	searchURL = "https://www.googleapis.com/youtube/v3/search"
)

type YouTubeResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
}

var ContentID []string
var ContentTitle []string
var ContentURL []string

func Search(query string, maxResults int) error {
	params := url.Values{}
	params.Add("part", "snippet")
	params.Add("maxResults", fmt.Sprintf("%d", maxResults))
	params.Add("q", query)
	params.Add("key", apiKey)

	apiURL := fmt.Sprintf("%s?%s", searchURL, params.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error searching videos, response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var ytResponse YouTubeResponse
	err = json.Unmarshal(body, &ytResponse)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	ContentID = make([]string, len(ytResponse.Items))
	ContentTitle = make([]string, len(ytResponse.Items))
	ContentURL = make([]string, len(ytResponse.Items))

	for i, item := range ytResponse.Items {
		ContentID[i] = item.ID.VideoID
		ContentTitle[i] = item.Snippet.Title
		ContentURL[i] = fmt.Sprintf("https://youtu.be/%s", item.ID.VideoID)
		fmt.Println(ContentTitle[i])
	}
	return nil
}
