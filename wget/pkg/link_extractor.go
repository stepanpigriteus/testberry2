package pkg

import (
	"bytes"

	"golang.org/x/net/html"
)

func LinkExtr(resp []byte) ([]string, error) {
	doc, err := html.Parse(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "a":
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			case "link", "script", "img":
				for _, attr := range n.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						links = append(links, attr.Val)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links, nil
}
