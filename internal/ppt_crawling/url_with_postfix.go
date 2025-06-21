package ppt_crawling

import (
	"fmt"

	"golang.org/x/net/html"
)

func URLwithPostfix(blocks []*html.Node, prefix string) ([]string, error) {

	var urls []string
	if len(blocks) == 0 {
		return nil, nil
	}

	for _, block := range blocks {
		hrefs, _ := getURLFromBlock(block)
		for _, href := range hrefs {

			urls = append(urls, fmt.Sprintf("%s%s", href, prefix))
		}
	}

	if len(urls) == 0 {
		fmt.Println("no download url found in URLwithPostfix")
		return nil, nil
	}
	return urls, nil
}
