package yandexmusic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type YandexMusicClient struct {
	BaseURL    string
	OAuthToken string
}

func NewYandexMusicClient(token string) *YandexMusicClient {
	return &YandexMusicClient{
		BaseURL:    "https://api.music.yandex.net",
		OAuthToken: token,
	}
}

// AccountStatus возвращает текущий трек пользователя
func (c *YandexMusicClient) AccountStatus() (*Account, error) {
	url := c.BaseURL + "/account/status"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "OAuth "+c.OAuthToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch account status")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Result GetAccountResult `json:"result"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result.Result.Account, nil
}
