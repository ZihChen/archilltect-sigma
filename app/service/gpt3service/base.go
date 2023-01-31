package gpt3service

import (
	"archilltect-sigma/app/settings"
	"archilltect-sigma/app/structs"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type Interface interface {
	Completion(ctx context.Context, request structs.CompletionRequest) (response structs.CompletionResponse, err error)
}

type service struct {
	httpClient *http.Client
	model      Model
	baseURL    string
	apiKey     string
}

var singleton *service
var once sync.Once

func New(options ...Option) Interface {
	ser := &service{
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
		apiKey:  os.Getenv("GPT_KEY"),
		baseURL: settings.Config.GptConfig.BaseUrl,
		model:   DavinciModel,
	}

	for _, opt := range options {
		_ = opt(ser)
	}

	once.Do(func() {
		singleton = ser
	})
	return singleton
}

func (s *service) Completion(ctx context.Context, request structs.CompletionRequest) (response structs.CompletionResponse, err error) {
	request.Model = string(s.model)
	req, err := s.newRequest(ctx, "POST", "/completions", request)
	if err != nil {
		return
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		zap.L().Error("[Send request failed]:", zap.Any("error", err))
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errResp := &structs.ErrorResponse{}
		data, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(data, &errResp)
		zap.L().Error("[Gpt3 http error]:", zap.Any("status code", resp.StatusCode), zap.Any("error msg", errResp.Error))
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		zap.L().Error("[Body decode error]:", zap.Any("error", err))
		return
	}
	return
}

func (s *service) newRequest(ctx context.Context, method, path string, payload interface{}) (*http.Request, error) {
	raw, err := json.Marshal(payload)
	if err != nil {
		zap.L().Error("[Payload marshal error]:", zap.Any("error", err))
		return nil, err
	}
	url := s.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(raw))
	if err != nil {
		zap.L().Error("[New request error]:", zap.Any("error", err))
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
	return req, nil
}
