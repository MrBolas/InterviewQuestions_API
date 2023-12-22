package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type OllamaClient struct {
	baseurl string
}

type PromptRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type PromptResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

func NewOllamaClient() *OllamaClient {
	return &OllamaClient{
		baseurl: "http://127.0.0.1:11434",
	}
}

func (c *OllamaClient) GenerateQuestion(lang string, category string, level string) (string, error) {

	questionPrompt := fmt.Sprintf("Generate an %s 45 minutes programming exercise for an interview. This exercise should test the candidate knowledge in data structures in language: %s. The question should approach topics in the %s category. The generated exercise should have a list of goals and a few additional bonus goals.", level, lang, category)
	return c.Prompt(questionPrompt).Response, nil
}

func (c *OllamaClient) GenerateAnswer(question string) (string, error) {

	responsePrompt := fmt.Sprintf("Give me an answer for the following question, generate code example: %s", question)
	return c.Prompt(responsePrompt).Response, nil
}

func (c *OllamaClient) Prompt(prompt string) *PromptResponse {

	newPrompt := PromptRequest{
		Model:  "Mistral",
		Prompt: prompt,
		Stream: false,
	}

	jsonStr, err := json.Marshal(newPrompt)
	if err != nil {
		log.Printf("Error marshaling request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Post(c.baseurl+"/api/generate", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var data PromptResponse
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)

	return &data
}
