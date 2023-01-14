package kogpt

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneration(t *testing.T) {
	KAKAO_API_KEY, ok := os.LookupEnv("KAKAO_API_KEY")
	if !ok {
		t.Skip("KAKAO_API_KEY is not set")
	}

	c := NewClient(
		http.DefaultClient,
		KAKAO_API_KEY,
	)

	// https://developers.kakao.com/docs/latest/ko/kogpt/rest-api#request-sample
	result, err := c.Generation(context.Background(), GenerationParams{
		Prompt:    "오늘 아침 하늘은 곧 비가 올 것 같아서",
		MaxTokens: 120,
		N:         2,
	})

	assert.NoError(t, err)
	assert.Len(t, result.Generations, 2)
	assert.Greater(t, len(result.Generations[0].Text), 10)
	assert.Greater(t, len(result.Generations[1].Text), 10)
}
