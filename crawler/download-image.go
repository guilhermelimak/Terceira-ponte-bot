package crawler

import (
	"io"
	"net/http"
	"os"
)

// SaveImage : Save image from url to disk
func SaveImage(url string, path string) {
	img, _ := os.Create(path)
	defer img.Close()

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	_, err := io.Copy(img, resp.Body)
	if err != nil {
		panic(err)
	}
}
