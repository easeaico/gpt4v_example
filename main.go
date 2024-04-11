package main

import (
	"flag"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/imgo"
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

	img := robotgo.CaptureImg(10, 10, 300, 300)
	imgo.Save("weixin.png", img)

	/*
		gpt4Client := NewGpt4Client(openaiToken)
			resp, err := gpt4Client.AskGPT4V(img, chatMsg)
			if err != nil {
				fmt.Printf("ask gpt4 error: %v\n", err)
			} else {
				fmt.Println(resp)
			}
	*/
}
