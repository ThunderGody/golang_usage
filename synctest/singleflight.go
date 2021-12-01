package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	g            singleflight.Group
	ErrCacheMiss = errors.New("cache miss")
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := load("key")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()

	}
	wg.Wait()
}

func load(key string) (string, error) {
	data, err := loadFromCache(key)
	if err != nil && err == ErrCacheMiss {
		v, err, _ := g.Do(key, func() (interface{}, error) {
			data, err := loadFromDB(key)
			if err != nil {
				return nil, err
			}
			setCache(key, data)
			return data, nil
		})
		if err != nil {
			log.Println(err)
			return "", err
		}
		data = v.(string)
	}

	return data, err
}
func loadFromCache(key string) (string, error) {
	return "", ErrCacheMiss
}

func setCache(key, data string) {}

func loadFromDB(key string) (string, error) {
	fmt.Println("query db")
	unix := strconv.Itoa(int(time.Now().UnixNano()))
	return unix, nil
}
