package kimi4go

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

func KimiAi(question string, roleSystemContent string) (string, error) {
	ctx := context.Background()

	client := NewClient[moonshot](moonshot{
		baseUrl: "https://api.moonshot.cn/v1",
		key:     os.Getenv("MOONSHOT_API_KEY"),
		client:  http.DefaultClient,
		log: func(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration) {
			log.Printf("[%s] %s %s", caller, request.URL, elapse)
		},
	})

	completion, err := client.CreateChatCompletion(ctx, &ChatCompletionRequest{
		Messages: []*Message{
			{
				Role:    RoleSystem,
				Content: &Content{Text: roleSystemContent},
			},
			{
				Role:    RoleUser,
				Content: &Content{Text: question},
			},
		},
		Model:       ModelMoonshot8K,
		MaxTokens:   4096,
		N:           1,
		Temperature: "0.3",
	})

	if err != nil {
		return "", err
	}
	return completion.GetMessageContent(), nil
}
