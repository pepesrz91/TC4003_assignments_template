package cos418_hw1_1

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
// 	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
	// HINT: To keep only alphanumeric characters, use the regex "[^0-9a-zA-Z]+"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	// Split  the data and format it correctly
	fileString := string(data)
	toLowerCase := strings.ToLower(fileString)
	noQuotation := strings.Replace(toLowerCase, "'", "", -1)
	regularExp := regexp.MustCompile("[^0-9a-zA-Z]+")
	fileFiltered := regularExp.Split(noQuotation, -1)

	// noDots := strings.Replace(toLowerCase, ".", "", -1)
	// noCommas := strings.Replace(noDots, ",", "", -1)
	// noQuotation := strings.Replace(noCommas, "'", "", -1)
	// stringArray := strings.Fields(noQuotation)

	// Do a for loop to check wether a word is greater or euqal  to  charThreshold
	wordsFrequency := make(map[string]int)

	for _, item := range fileFiltered {
		_, itemMatched := wordsFrequency[item]
		if itemMatched {
			wordsFrequency[item]++
		} else {
			wordsFrequency[item] = 1
		}
	}

	// fmt.Printf("%v \n", wordsFrequency)
	wordsThreshold := make(map[string]int)
	for word, item := range wordsFrequency {
		if len(word) >= charThreshold {
			wordsThreshold[word] = item
		} else {
			continue
		}
	}

	// Append
	var ss []WordCount
	for k, v := range wordsThreshold {
		ss = append(ss, WordCount{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Count > ss[j].Count
	})

	count := 0
	//topWords := make(map[string]int)

	sortWordCounts(ss)

	var topWordCount []WordCount
	for _, kv := range ss {
		if count >= numWords {
			break
		}
		topWordCount = append(topWordCount, kv)
		count++
	}

	//fmt.Printf("Top words %v", ss)
	// keys := make([]string, 0, len(wo))
	// for k := range wordsThreshold {
	// 	keys = append(keys, k)
	// }
	// sort.Strings(keys)

	// for _, k := range keys {
	// 	fmt.Println(k, m[k])
	// }

	// var topWords []string

	// for k := range wordsThreshold {
	// 	topWords = append(topWords, k)
	// }

	// sort.Strings(topWords)

	// fmt.Printf("%v \n", topWords)

	// pl := make(PairList, len(wordFrequencies))
	// i := 0
	// for k, v := range wordFrequencies {
	// 	pl[i] = Pair{k, v}
	// 	i++
	// }
	// sort.Sort(sort.Reverse(pl))

	return topWordCount
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
