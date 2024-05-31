package fediverse

import (
	"fmt"
	"net/http"
	"strings"
)

const postPath = "/api/v1/statuses"

type Fediverse struct {
	apiKey string
	url    string
	client http.Client
}

func New(apiKey string, url string) *Fediverse {
	return &Fediverse{apiKey, url, http.Client{}}
}

func (f *Fediverse) Post(message string) error {
	req, err := http.NewRequest(
		http.MethodPost,
		f.url+postPath,
		strings.NewReader(fmt.Sprintf("status=%v", message)),
	)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", f.apiKey))

	resp, err := f.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
