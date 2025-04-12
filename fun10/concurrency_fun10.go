package fun10

type (
	WebsiteChecker func(string) bool
	result         struct {
		string
		bool
	}
)

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}

func ConcCheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			// Send Statement
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		// Receive Statement
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
