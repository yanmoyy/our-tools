package main

import (
	"fmt"
	"html"
	"strings"
)

func commandNum(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: title <query>")
	}
	query := args[:]
	blank := html.EscapeString(" ")
	fullQuery := strings.Join(query, blank)

	fmt.Println("fullQuery: ", fullQuery)

	response, err := cfg.client.GoogleSearch(fullQuery)
	if err != nil {
		return err
	}

	println("start")

	var link string
	for _, item := range response.Items {
		if strings.HasPrefix(item.Link, TistoryNum) {
			link = item.Link
		}
	}
	fmt.Println(link)
	println("end")

	return nil
}
