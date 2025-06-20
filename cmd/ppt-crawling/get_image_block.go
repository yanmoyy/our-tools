package main

import (
	"fmt"
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
	fmt.Printf("imagesBlock size : %d\n", len(imageblocks))
	return imageblocks, nil
}
