package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GPTChatRequest struct {
	Model    string       `json:"model"`
	Messages []GPTMessage `json:"messages"`
}

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTResponse struct {
	Choices []GPTChoice `json:"choices"`
}

type GPTChoice struct {
	Message GPTMessage `json:"message"`
}

func chatChatGPT(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	data := GPTChatRequest{
		Model: "gpt-4-turbo-preview",
		Messages: []GPTMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}
	bytesRepresentation, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var gptResponse GPTResponse
	err = json.Unmarshal(body, &gptResponse)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return gptResponse.Choices[0].Message.Content, nil
}
