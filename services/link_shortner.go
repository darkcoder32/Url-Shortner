package services

import (
	"errors"
	"fmt"
	"gin/shorti/myconfig"
	"strings"

	"github.com/dchest/uniuri"
)

func GetLongUrl(shortUrl string) (string, error) {
	longUrl, err := myconfig.RedisClient.Get(shortUrl).Result()
	if err != nil {
		fmt.Printf("Error in fetching value from redis: %v\n", err)
	}
	if longUrl != "" {
		fmt.Printf("Long Url found in redis: %v\n", longUrl)
		return longUrl, nil
	}
	r, ok := myconfig.MyMap[shortUrl]
	if !ok {
		return "", errors.New("short url not found")
	}
	myconfig.RedisClient.Set(shortUrl, r, 0)
	return r, nil
}

func CreateShortUrl(longUrl string) (string, error) {
	for _, v := range myconfig.MyMap {
		if strings.EqualFold(longUrl, v) {
			return v, nil
		}
	}
	randomUniqueId := uniuri.NewLen(7)
	myconfig.MyMap[randomUniqueId] = longUrl
	return randomUniqueId, nil
}
