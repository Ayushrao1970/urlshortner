package main

import (
	"crypto/md5"
	"fmt"
	"time"
	
)

type url struct {
	Id           string    `json:"id"`
	OriginalUrl  string    `json:"originalurl"`
	ShortUrl     string    `json:"shorturl"`
	Creationdate time.Time `json:"creationdate"`
}
var urlDB = make (map[string]url)
func genererateShortUrl(OriginalUrl string) string{
hasher:=md5.New()
hasher.Write([]byte(OriginalUrl))
fmt.Println("hasher",hasher)
return"https://www.youtube.com/watch?v=dVVJU-3eU1g&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=34"


}
func main() {
	OriginalUrl:="https://www.youtube.com/watch?v=dVVJU-3eU1g&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=34"
	genererateShortUrl(OriginalUrl)
	
}
