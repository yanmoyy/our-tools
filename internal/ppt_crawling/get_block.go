package ppt_crawling

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func getBlocks(htmlbody, nodeType, key, val string) ([]*html.Node, error) {
	htmlReader := strings.NewReader(htmlbody)

	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("error Parse in getBlocks\n%v", err)
	}

	blocks := traverseNodes(doc, nodeType, key, val)
	return blocks, nil
}
