package ppt_crawling

import (
	"fmt"
)

func crawlDownloadURLs(htmlbody, nodeType, key, val, prefix string) ([]string, error) {

	blocks, err := getBlocks(htmlbody, nodeType, key, val)
	if err != nil {
		return nil, fmt.Errorf("error get imageblock in getDownloadURLs: %v", err)
	}

	urls, err := URLwithPostfix(blocks, prefix)
	if err != nil {
		return nil, fmt.Errorf("error append postfix string to imageblock in getDownloadURLs\n%v", err)
	}
	return urls, nil

}
