package ppt_crawling

import (
	"fmt"
)

func getDownloadUrls(response googleSearchResponse) ([]string, error) {
	links := getSiteLink(response)

	var urls []string
	count := 0
	for _, link := range links {
		if link == "" {
			continue
		}
		htmlbody, err := getHTML(link)
		if err != nil {
			fmt.Printf("error get HTML in Search\n%v\n", err)
			continue
		}
		imageUrls, imageErr := crawlDownloadURLs(htmlbody, "span", "class", "imageblock", "?original")
		fileUrls, fileErr := crawlDownloadURLs(htmlbody, "figure", "class", "fileblock", "")

		if imageErr != nil && fileErr != nil {
			return nil, fmt.Errorf("error get download urls (fileURLs) in Search\n%v", err)
		}
		url := filterUrls(imageUrls, fileUrls)
		if url == "" {
			continue
		}
		count++
		urls = append(urls, url)
	}
	fmt.Printf("Found %d urls\n", count)
	fmt.Println()

	return urls, nil

}

func getSiteLink(response googleSearchResponse) []string {
	var links []string
	for _, item := range response.Items {

		links = append(links, item.Link)
	}
	if len(links) == 0 {
		fmt.Println("no link found in getSiteLink")
		fmt.Println()
		fmt.Println()
		return nil
	}

	return links
}
