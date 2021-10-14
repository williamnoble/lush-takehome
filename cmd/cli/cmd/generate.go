package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a shortURL",
	Long:  `Generate a short URL`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		fmt.Println(args[0])
		url := args[0]
		generateShortURL(url)

	},
}

type Input struct {
	Url string `json:"url"`
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateShortURL(u string) {

	// Override default client with timeout
	client := &http.Client{Timeout: 10 * time.Second}

	input := Input{
		Url: u,
	}

	js, err := json.Marshal(input)
	if err != nil {
		fmt.Println("failed to marshal input ", err)
		return
	}

	b := bytes.NewReader(js)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/shorten", b)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("encountered an error when client performed http request ", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed to read the response body ", err)
		return
	}
	fmt.Println(string(body))
}
