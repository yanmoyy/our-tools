package ppt_crawling

import (
	"golang.org/x/net/html"
)

func traverseNodes(node *html.Node, data, key, val string) []*html.Node {
	var nodes []*html.Node

	if node.Type == html.ElementNode && node.Data == data {
		for _, anchor := range node.Attr {
			if anchor.Key == key && anchor.Val == val {
				nodes = append(nodes, node)
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		nodes = append(nodes, traverseNodes(child, data, key, val)...)
	}
	return nodes
}
func getDataFromKey(node *html.Node, data, key string) []string {
	var urls []string

	if node.Type == html.ElementNode && node.Data == data {
		for _, anchor := range node.Attr {
			if anchor.Key == key {
				urls = append(urls, anchor.Val)

			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		urls = append(urls, getDataFromKey(child, data, key)...)
	}
	return urls
}
