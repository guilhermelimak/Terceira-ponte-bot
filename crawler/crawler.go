package crawler

import (
	"net/http"
	"regexp"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// GetImgLinks : Get all camera image html elements from URL
func GetImgLinks(url string) []*html.Node {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	fileName := regexp.MustCompile(`.*camera[0-9]+TP.*`)

	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.A && scrape.Attr(n, `class`) == `imagem` {
			return fileName.MatchString(n.Attr[0].Val)
		}

		return false
	}

	return scrape.FindAll(root, matcher)
}
