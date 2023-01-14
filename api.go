package kogpt

import (
	"context"
)

// https://developers.kakao.com/docs/latest/ko/kogpt/rest-api#request-request
type GenerationParams struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	N           int     `json:"n,omitempty"`
}

// https://developers.kakao.com/docs/latest/ko/kogpt/rest-api#request-response
type GenerationResult struct {
	Id          string       `json:"id"`
	Generations []Generation `json:"generations"`
	Usage       Usage        `json:"usage"`
}

// https://developers.kakao.com/docs/latest/ko/kogpt/rest-api#request-response-generation
type Generation struct {
	Text   string `json:"text"`
	Tokens int    `json:"tokens"`
}

// https://developers.kakao.com/docs/latest/ko/kogpt/rest-api#request-response-usage
type Usage struct {
	PromptTokens    int `json:"prompt_tokens"`
	GeneratedTokens int `json:"generated_tokens"`
	TotalTokens     int `json:"total_tokens"`
}

func (c *Client) Generation(ctx context.Context, params GenerationParams) (*GenerationResult, error) {
	result := &GenerationResult{}
	err := c.issueRequest(ctx, "generation", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
