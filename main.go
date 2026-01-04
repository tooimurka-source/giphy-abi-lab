package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GiphyResponse struct {
	Data []Gif `json:"data"`
}

type Gif struct {
	Title  string `json:"title"`
	Images Images `json:"images"`
}

type Images struct {
	Original Original `json:"original"`
}

type Original struct {
	URL string `json:"url"`
}

func getGifs(apiKey, query string, limit int) (*GiphyResponse, error) {
	baseURL := "https://api.giphy.com/v1/gifs/search"

	params := url.Values{}
	params.Add("api_key", apiKey)
	params.Add("q", query)
	params.Add("limit", fmt.Sprint(limit))

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GiphyResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	var keyword string
	var limit int

	fmt.Print("Введіть ключове слово: ")
	fmt.Scanln(&keyword)

	fmt.Print("Кількість GIF: ")
	fmt.Scanln(&limit)

	apiKey := "GiIcanazu0FO8ait6rNpgnab3SMXiHOA"

	response, err := getGifs(apiKey, keyword, limit)
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}

	for i, gif := range response.Data {
		fmt.Printf("\n%d) Назва: %s\nПосилання: %s\n",
			i+1, gif.Title, gif.Images.Original.URL)
	}
}
