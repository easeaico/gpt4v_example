package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"image"
	"image/png"
	"io"

	"github.com/sashabaranov/go-openai"
)

var (
	openaiToken string
	chatMsg     string
)

func init() {
	flag.StringVar(&openaiToken, "t", "", "Openai API Token")
	flag.StringVar(&chatMsg, "m", "", "chat message to describe task")
}

func main() {
	flag.Parse()

	ctl := NewScreenController(0, image.Rect(0, 25, 1050, 1100))
	img, err := ctl.TakeScreenshot()
	if err != nil {
		fmt.Printf("take screenshot error: %v\n", err)
		return
	}
	var b bytes.Buffer
	png.Encode(&b, img)
	data, err := io.ReadAll(&b)
	if err != nil {
		fmt.Printf("read image data error: %v\n", err)
		return
	}

	e64 := base64.StdEncoding
	imageData := e64.EncodeToString(data)
	urlData := fmt.Sprintf("data:image/png;base64,%s", imageData)

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
