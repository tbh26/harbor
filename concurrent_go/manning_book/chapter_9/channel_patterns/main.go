package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

const downloaders = 20

func generateUrls(quit <-chan bool) <-chan string {
	urls := make(chan string)
	go func() {
		defer close(urls)
		//for i := 1000; i <= 1030; i++ {
		for i := 100; i <= 130; i++ {
			url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
			time.Sleep(12 * time.Millisecond)
			select {
			case urls <- url:
			case q := <-quit:
				fmt.Printf("Quitting urls gen, received: '%t' \n", q)
				return
			}
		}
	}()
	return urls
}

func downloadPages(quit <-chan bool, urls <-chan string) <-chan string {
	pages := make(chan string)
	go func() {
		defer close(pages)
		moreData, url := true, ""
		for moreData {
			select {
			case url, moreData = <-urls:
				if moreData {
					resp, _ := http.Get(url)
					if resp.StatusCode != 200 {
						panic("Server’s error: " + resp.Status)
					}
					body, _ := io.ReadAll(resp.Body)
					pages <- string(body)
					_ = resp.Body.Close()
					fmt.Printf(" - %q downloaded (%d) \n", url, len(body))
				}
			case q := <-quit:
				fmt.Printf("Quitting download pages, received: '%t' \n", q)
				return
			}
		}
		// fmt.Println()
	}()
	return pages
}

func longestWords(quit <-chan bool, words <-chan string) <-chan string {
	longWords := make(chan string)
	go func() {
		defer close(longWords)
		uniqueWordsMap := make(map[string]bool)
		uniqueWords := make([]string, 0)
		moreData, word := true, ""
		for moreData {
			select {
			case word, moreData = <-words:
				if moreData && !uniqueWordsMap[word] {
					uniqueWordsMap[word] = true
					uniqueWords = append(uniqueWords, word)
				}
			case <-quit:
				return
			}
		}
		sort.Slice(uniqueWords, func(a, b int) bool {
			return len(uniqueWords[a]) > len(uniqueWords[b])
		})
		longWords <- strings.Join(uniqueWords[:10], ", ")
	}()
	return longWords
}

func extractWords(quit <-chan bool, pages <-chan string) <-chan string {
	words := make(chan string)
	go func() {
		defer close(words)
		wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
		moreData, pg := true, ""
		for moreData {
			select {
			case pg, moreData = <-pages:
				if moreData {
					for _, word := range wordRegex.FindAllString(pg, -1) {
						words <- strings.ToLower(word)
					}
				}
			case q := <-quit:
				fmt.Printf("Quitting extract words, received: '%t' \n", q)
				return
			}
		}
	}()
	return words
}

func fanIn[K any](quit <-chan bool, allChannels ...<-chan K) chan K {
	wg := sync.WaitGroup{}
	wg.Add(len(allChannels))
	output := make(chan K)
	for _, c := range allChannels {
		go func(channel <-chan K) {
			defer wg.Done()
			for i := range channel {
				select {
				case output <- i:
				case <-quit:
					return
				}
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

func fanOutIn() {
	quit := make(chan bool)
	defer close(quit)
	urls := generateUrls(quit)
	pages := make([]<-chan string, downloaders)
	for i := 0; i < downloaders; i++ {
		pages[i] = downloadPages(quit, urls)
	}
	results := longestWords(quit, extractWords(quit, fanIn(quit, pages...)))
	fmt.Println("\nLongest Words:", <-results)
}

func CreateAll[K any](n int) []chan K {
	channels := make([]chan K, n)
	for i, _ := range channels {
		channels[i] = make(chan K)
	}
	return channels
}

func CloseAll[K any](channels ...chan K) {
	for _, output := range channels {
		close(output)
	}
}

func broadcast[K any](quit <-chan bool, input <-chan K, n int) []chan K {
	outputs := CreateAll[K](n)
	go func() {
		defer CloseAll(outputs...)
		var msg K
		moreData := true
		for moreData {
			select {
			case msg, moreData = <-input:
				if moreData {
					for _, output := range outputs {
						output <- msg
					}
				}
			case <-quit:
				return
			}
		}
	}()
	return outputs
}

func frequentWords(quit <-chan bool, words <-chan string) <-chan string {
	mostFrequentWords := make(chan string)
	go func() {
		defer close(mostFrequentWords)
		freqMap := make(map[string]int)
		freqList := make([]string, 0)
		moreData, word := true, ""
		for moreData {
			select {
			case word, moreData = <-words:
				if moreData {
					if freqMap[word] == 0 {
						freqList = append(freqList, word)
					}
					freqMap[word] += 1
				}
			case <-quit:
				return
			}
		}
		sort.Slice(freqList, func(a, b int) bool {
			return freqMap[freqList[a]] > freqMap[freqList[b]]
		})
		mostFrequentWords <- strings.Join(freqList[:10], ", ")
	}()
	return mostFrequentWords
}

func useBroadcast() {
	quit := make(chan bool)
	defer close(quit)
	urls := generateUrls(quit)
	pages := make([]<-chan string, downloaders)
	for i := 0; i < downloaders; i++ {
		pages[i] = downloadPages(quit, urls)
	}
	words := extractWords(quit, fanIn(quit, pages...))
	wordsMulti := broadcast(quit, words, 2)
	longestResults := longestWords(quit, wordsMulti[0])
	frequentResults := frequentWords(quit, wordsMulti[1])
	fmt.Println("\nLongest Words:", <-longestResults)
	fmt.Println("\nMost frequent Words:", <-frequentResults)
}

func take[K any](quit chan bool, n int, input <-chan K) <-chan K {
	output := make(chan K)
	go func() {
		defer close(output)
		moreData := true
		var msg K
		for n > 0 && moreData {
			select {
			case msg, moreData = <-input:
				if moreData {
					output <- msg
					n--
				}
			case <-quit:
				return
			}
		}
		if n == 0 {
			close(quit)
		}
	}()
	return output
}

func takeNdemo() {
	quitWords := make(chan bool)
	quit := make(chan bool)
	defer close(quit)
	urls := generateUrls(quitWords)
	pages := make([]<-chan string, downloaders)
	for i := 0; i < downloaders; i++ {
		pages[i] = downloadPages(quitWords, urls)
	}
	words := take(quitWords, 10000, extractWords(quitWords, fanIn(quitWords, pages...)))
	wordsMulti := broadcast(quit, words, 2)
	longestResults := longestWords(quit, wordsMulti[0])
	frequentResults := frequentWords(quit, wordsMulti[1])

	fmt.Println("Longest Words:", <-longestResults)
	fmt.Println("Most frequent Words:", <-frequentResults)
}

func primeMultipleFilter(numbers <-chan int, quit chan<- bool) {
	var right chan int
	p := <-numbers
	fmt.Printf("%d ", p)
	for n := range numbers {
		if n%p != 0 {
			if right == nil {
				right = make(chan int)
				go primeMultipleFilter(right, quit)
			}
			right <- n
		}
	}
	if right == nil {
		close(quit)
	} else {
		close(right)
	}
}

func primeDemo() {
	numbers := make(chan int)
	quit := make(chan bool)
	go primeMultipleFilter(numbers, quit)
	for i := 2; i < 100000; i++ {
		numbers <- i
	}
	close(numbers)
	<-quit
}

func demo() {
	fmt.Println("\n\t\t=-= fan out / in =-=\n")
	fanOutIn()
	fmt.Println("\n\t\t=-= use broadcast =-=\n")
	useBroadcast()
	fmt.Println("\n\t\t=-= take n =-=\n")
	takeNdemo()
	fmt.Println("\n\t\t=-= prime (with channels) =-=\n")
	primeDemo()
	fmt.Println("\n\t\t=-= done =-=\n")
}

func main() {
	fmt.Println(" == Channel programming II ⛓️  ==")
	fmt.Println()

	demo()
}
