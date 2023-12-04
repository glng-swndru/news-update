package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
)

func main() {
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		fmt.Println("API key is not set")
		return
	}

	url := fmt.Sprintf("https://newsdata.io/api/1/news?apikey=%s", apiKey)

	statusCode, body, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	if statusCode != fasthttp.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", statusCode)
		return
	}

	news := gjson.Get(string(body), "results.#.title")
	news.ForEach(func(key, value gjson.Result) bool {
		fmt.Println(value.String())
		return true
	})

	file, err := os.Create("README.md")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "# Latest News\n")
	news.ForEach(func(key, value gjson.Result) bool {
		fmt.Fprintln(file, "- ", value.String())
		return true
	})

	fmt.Println("README.md updated successfully.")
}

func init() {
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc
}
