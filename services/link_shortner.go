package services

import (
	"errors"
	"gin/shorti/myconfig"
	"strings"

	"github.com/dchest/uniuri"
)

func GetLongUrl(shortUrl string) (string, error) {
	r, ok := myconfig.MyMap[shortUrl]
	if !ok {
		return "", errors.New("short url not found")
	}
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
