package crawling_api

import (
	"strings"

	"golang.org/x/net/html"
)

func getImageBlocks(htmlbody string) ([]*html.Node, error) {
	htmlReader := strings.NewReader(htmlbody)

	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	imageblocks := traverseNodes(doc, "span", "class", "imageblock")
	return imageblocks, nil
}

func getFileBlocks(htmlbody string) ([]*html.Node, error) {
	htmlReader := strings.NewReader(htmlbody)

	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	fileblocks := traverseNodes(doc, "figure", "class", "fileblock")
	return fileblocks, nil
}
