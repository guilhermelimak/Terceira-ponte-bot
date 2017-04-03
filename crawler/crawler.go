package crawler

import (
	"net/http"
	"regexp"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// GetPage : Get all camera image html elements from URL
func GetPage(url string) []*html.Node {
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
		if len(n.Attr) > 0 && n.DataAtom == atom.A {
			return fileName.MatchString(n.Attr[0].Val)
		}

		return false
	}

	return scrape.FindAll(root, matcher)
}
