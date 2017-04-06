package crawler

import (
	"io"
	"net/http"
	"os"
)

// SaveImage : Save image from url to disk
func SaveImage(c chan int64, url string, path string) {
	img, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer img.Close()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, err := io.Copy(img, resp.Body)
	if err != nil {
		panic(err)
	}
	c <- result
}
