package concurrency

type WebsiteChecker func(string) bool

type result struct {
	webUrl  string
	isValid bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	//Do the processing in parallel
	for _, url := range urls {
		go func(u string) {
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	//Unpack the channel into the map
	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.webUrl] = r.isValid
	}

	return results
}
