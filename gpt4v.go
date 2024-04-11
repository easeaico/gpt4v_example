package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io"

	"github.com/sashabaranov/go-openai"
)

type GPT4Client struct {
	client *openai.Client
}

func NewGpt4Client(apiToken string) *GPT4Client {
	return &GPT4Client{
		client: openai.NewClient(openaiToken),
	}
}

func (c *GPT4Client) AskGPT4V(img image.Image, chatMsg string) (string, error) {
	var b bytes.Buffer
	png.Encode(&b, img)
	data, err := io.ReadAll(&b)
	if err != nil {
		return "", fmt.Errorf("read image data error: %w", err)
	}

	e64 := base64.StdEncoding
	imageData := e64.EncodeToString(data)
	urlData := fmt.Sprintf("data:image/png;base64,%s", imageData)

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4VisionPreview,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					MultiContent: []openai.ChatMessagePart{
						{
							Type: openai.ChatMessagePartTypeText,
							Text: chatMsg,
						},
						{
							Type: openai.ChatMessagePartTypeImageURL,
							ImageURL: &openai.ChatMessageImageURL{
								URL: urlData,
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("creatingChatCompletion error: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}
