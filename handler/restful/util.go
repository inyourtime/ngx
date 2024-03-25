package restful

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"ngx/domain"
)

type GitHubUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GithubUserEmail struct {
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	Primary  bool   `json:"primary"`
}

func GetGoogleUserInfo(endpoint string) (*domain.GoogleUserInfo, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	info := domain.GoogleUserInfo{}
	if err := json.Unmarshal(userData, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func GetGithubUserInfo(endpoint string, access_token string) (*GitHubUser, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+access_token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	info := GitHubUser{}
	if err := json.Unmarshal(userData, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func GetGithubUserEmail(endpoint string, access_token string) (*[]GithubUserEmail, error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+access_token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	info := []GithubUserEmail{}
	if err := json.Unmarshal(userData, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func GenerateState() (string, error) {
	// Generate a random byte slice to use as the state value
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64 URL encoding
	state := base64.URLEncoding.EncodeToString(randomBytes)

	return state, nil
}
