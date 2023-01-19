package structs

type CompletionResponseChoice struct {
	Text         string        `json:"text"`
	Index        int           `json:"index"`
	LogProbs     LogprobResult `json:"logprobs"`
	FinishReason string        `json:"finish_reason"`
}

type LogprobResult struct {
	Tokens        []string             `json:"tokens"`
	TokenLogprobs []float32            `json:"token_logprobs"`
	TopLogprobs   []map[string]float32 `json:"top_logprobs"`
	TextOffset    []int                `json:"text_offset"`
}

type CompletionResponse struct {
	ID      string                     `json:"id"`
	Object  string                     `json:"object"`
	Created int                        `json:"created"`
	Model   string                     `json:"model"`
	Choices []CompletionResponseChoice `json:"choices"`
	Usage   CompletionResponseUsage    `json:"usage"`
}

type CompletionResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type CompletionRequest struct {
	// 模型種類
	Model string `json:"model"`
	// 輸入文字
	Prompt string `json:"prompt"`
	// 輸出字數
	MaxTokens int `json:"max_tokens"`
	// 產生結果數
	N int `json:"n"`
}

type ErrorResponse struct {
	Error APIError `json:"error"`
}

type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
