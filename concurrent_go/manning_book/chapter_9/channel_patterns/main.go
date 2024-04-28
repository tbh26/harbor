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

//var randSrc = rand.NewSource(time.Now().UnixNano())
//var random = rand.New(randSrc)
//
//func generateRandomWord(strLen int) string {
//	chars := "abcdefghijklmnopqrstuvwxyz"
//	word := make([]byte, strLen)
//	for i := range word {
//		word[i] = chars[random.Intn(len(chars))]
//	}
//	return string(word)
//}
//
//func wordsReceiver(words <-chan string, quit chan bool, max int) {
//	go func() {
//		var word string
//		fmt.Printf("wordsReceiver, words: '%v', quit: '%v', max: %d \n", words, quit, max)
//		for i := 0; i < max; i++ {
//			word = <-words
//			fmt.Printf(" - received: %q  (%d)\n", word, len(word))
//			time.Sleep(123 * time.Millisecond)
//		}
//		close(quit)
//	}()
//}

//func firstDemo() {
//	words := make(chan string)
//	quit := make(chan bool)
//	wordsReceiver(words, quit, 5)
//	var next string
//	nextLen := 2
//	for {
//		next = generateRandomWord(nextLen)
//		nextLen += 1
//		select {
//		case words <- next:
//		case q := <-quit:
//			fmt.Printf("Quitting words gen. (%v) \n", q)
//			return
//		}
//	}
//}

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

//func nextDemo() {
//	quit := make(chan bool)
//	defer close(quit)
//	results := generateUrls(quit)
//	for result := range results {
//		fmt.Printf("Result: %q \n", result)
//		if result == "https://www.rfc-editor.org/rfc/rfc1025.txt" {
//			quit <- true
//		}
//	}
//	fmt.Println("next demo done.")
//}

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

//func thirdDemo() {
//	quit := make(chan bool)
//	defer close(quit)
//	results := downloadPages(quit, generateUrls(quit))
//	for result := range results {
//		fmt.Printf(" - result length: %d \n", len(result))
//	}
//	fmt.Println("third demo done.")
//}

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

func demo() {
	fmt.Println("\n\t\t=-= fan out / in =-=\n")
	fanOutIn()
	fmt.Println("\n\t\t=-= use broadcast =-=\n")
	useBroadcast()
	fmt.Println("\n\t\t=-= done =-=\n")
}

func main() {
	fmt.Println(" == Channel programming II ⛓️  ==")
	fmt.Println()

	demo()
}
