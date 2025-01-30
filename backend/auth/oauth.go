package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type OAuthProvider struct {
	Config *oauth2.Config
	Name   string
}

func NewGithubProvider() *OAuthProvider {
	return &OAuthProvider{
		Name: "github",
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		},
	}
}

func NewGoogleProvider() *OAuthProvider {
	return &OAuthProvider{
		Name: "google",
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

type UserInfo struct {
	ID        string
	Email     string
	Name      string
	AvatarURL string
}

func (p *OAuthProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	client := p.Config.Client(oauth2.NoContext, token)

	var userInfo UserInfo
	var err error

	switch p.Name {
	case "github":
		userInfo, err = getGithubUserInfo(client)
	case "google":
		userInfo, err = getGoogleUserInfo(client)
	default:
		return nil, fmt.Errorf("unknown provider: %s", p.Name)
	}

	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func getGithubUserInfo(client *http.Client) (UserInfo, error) {
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return UserInfo{}, err
	}
	defer resp.Body.Close()

	var githubUser struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		Name      string `json:"name"`
		AvatarURL string `json:"avatar_url"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}, err
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
}

func getGoogleUserInfo(client *http.Client) (UserInfo, error) {
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return UserInfo{}, err
	}
	defer resp.Body.Close()

	var googleUser struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}, err
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
}
