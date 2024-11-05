package main

import (
	"fmt"
	encoder2 "url-shortener/encoder"
	"url-shortener/repository"
	"url-shortener/service"
)

func main() {
	encoder := encoder2.NewBase64Encoder("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	repo := repository.NewInMemoryURLRepository()

	service := service.NewURLShortenerSService(encoder, repo)

	shortURL, err := service.ShortenURL("www.google.com")
	if err != nil {
		fmt.Println("error shortening url ", err)
		return
	}

	fmt.Println("short url - ", shortURL)

	oldURL, err := service.Fetch(shortURL)
	if err != nil {
		fmt.Println("error fetching old url ", err)
		return
	}

	fmt.Println("old url - ", oldURL)
}
