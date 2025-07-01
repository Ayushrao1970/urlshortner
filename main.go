package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
	"net/http"
)

type URL struct {
	Id           string    `json:"id"`
	OriginalUrl  string    `json:"originalurl"`
	ShortUrl     string    `json:"shorturl"`
	CreationDate time.Time `json:"creationdate"`
}

var urlDB = make(map[string]URL)

func generateShortUrl(OriginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl))
	fmt.Println("hasher", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hashed data", data)
	hash := hex.EncodeToString(data)
	fmt.Println("encoded string", hash)
	fmt.Println("encoded string", hash[:8])
	return hash[:8]

}
func createUrl(originalUrl string) string {
	shortUrl := generateShortUrl(originalUrl)
	id := shortUrl
	urlDB[id] = URL{
		Id:           id,
		OriginalUrl:  originalUrl,
		ShortUrl:     shortUrl,
		CreationDate: time.Now()}
	return shortUrl

}
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("error not found")
	}
	return url, nil

}
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println("get method")
}
func main() {
	OriginalUrl := "https://www.youtube.com/watch?v=dVVJU-3eU1g&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=34"
	generateShortUrl(OriginalUrl)

	http.HandleFunc("/", handler)
fmt.Println("sever started")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting server:", err)
	}
}
