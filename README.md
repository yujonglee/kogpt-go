# KoGPT-Go
Unofficial Go SDK for KakaoBrain KoGPT

## Usage
```go
import (
	"context"
	"net/http"

	"github.com/yujong-lee/kogpt-go"
)

func Example() {
	c := kogpt.NewClient(
		http.DefaultClient,
		"KAKAO_API_KEY",
	)

	result, err := c.Generation(context.Background(), kogpt.GenerationParams{
		Prompt:    "오늘 아침 하늘은 곧 비가 올 것 같아서",
		MaxTokens: 120,
		N:         2,
	})
}
```
