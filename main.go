package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var (
	openaiToken string
	imageFile   string
	chatMsg     string
)

func init() {
	flag.StringVar(&openaiToken, "t", "", "Openai API Token")
	flag.StringVar(&imageFile, "f", "", "image file to upload")
	flag.StringVar(&chatMsg, "m", "", "chat message to describe task")
}

func main() {
	flag.Parse()

	f, err := os.Open(imageFile)
	if err != nil {
		fmt.Printf("open image file error: %v\n", err)
		return
	}

	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("read image data error: %v\n", err)
		return
	}

	e64 := base64.StdEncoding
	imageData := e64.EncodeToString(data)
	urlData := fmt.Sprintf("data:image/jpeg;base64,%s", imageData)

	client := openai.NewClient(openaiToken)
	resp, err := client.CreateChatCompletion(
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
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
