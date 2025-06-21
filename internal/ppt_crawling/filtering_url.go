package ppt_crawling

func filterUrls(imageUrls []string, fileUrls []string) string {
	var url string
	if imageUrls == nil {
		if fileUrls == nil {
			return ""
		}
		if len(fileUrls) == 3 {
			url = fileUrls[1]
		} else {
			url = fileUrls[0]
		}
	} else {
		if len(imageUrls) == 3 {
			url = imageUrls[1]
		} else {
			url = imageUrls[0]
		}
	}
	return url
}
