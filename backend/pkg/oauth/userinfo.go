package oauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type UserInfo struct {
	ID        string
	Email     string
	Name      string
	AvatarURL string
}

type providerConfig struct {
	userInfoURL string
	mapResponse func([]byte) (UserInfo, error)
}

var providerConfigs = map[string]providerConfig{
	"github": {
		userInfoURL: "https://api.github.com/user",
		mapResponse: func(body []byte) (UserInfo, error) {
			var githubUser struct {
				ID        int    `json:"id"`
				Email     string `json:"email"`
				Name      string `json:"name"`
				AvatarURL string `json:"avatar_url"`
			}
			if err := json.Unmarshal(body, &githubUser); err != nil {
				return UserInfo{}, err
			}
			return UserInfo{
				ID:        fmt.Sprintf("%d", githubUser.ID),
				Email:     githubUser.Email,
				Name:      githubUser.Name,
				AvatarURL: githubUser.AvatarURL,
			}, nil
		},
	},
	"google": {
		userInfoURL: "https://www.googleapis.com/oauth2/v2/userinfo",
		mapResponse: func(body []byte) (UserInfo, error) {
			var googleUser struct {
				ID      string `json:"id"`
				Email   string `json:"email"`
				Name    string `json:"name"`
				Picture string `json:"picture"`
			}
			if err := json.Unmarshal(body, &googleUser); err != nil {
				return UserInfo{}, err
			}
			return UserInfo{
				ID:        googleUser.ID,
				Email:     googleUser.Email,
				Name:      googleUser.Name,
				AvatarURL: googleUser.Picture,
			}, nil
		},
	},
}

func (p *Provider) getUserInfo(token *oauth2.Token) (*UserInfo, error) {
	config, ok := providerConfigs[p.name]
	if !ok {
		return nil, fmt.Errorf("unknown provider: %s", p.name)
	}

	client := p.config.Client(oauth2.NoContext, token)
	userInfo, err := fetchUserInfo(client, config)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func fetchUserInfo(client *http.Client, config providerConfig) (UserInfo, error) {
	resp, err := client.Get(config.userInfoURL)
	if err != nil {
		return UserInfo{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}, err
	}

	return config.mapResponse(body)
}
