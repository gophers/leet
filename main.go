package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	leetMap map[byte]string
)

func init() {
	leetMap = map[byte]string{
		'o': "0",
		'l': "1",
		'z': "2",
		'e': "3",
		'a': "4",
		's': "5",
		'b': "6",
	}
}

func main() {
	words, err := ioutil.ReadFile("words.txt")
	if err != nil {
		panic(err)
	}
	distFile, err := os.OpenFile("dist.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer distFile.Close()
	for _, word := range strings.Split(string(words), "\n") {
		if word == "" {
			continue
		}
		for _, lw := range leet(word) {
			fmt.Fprintf(distFile, "%s\n", lw)
			fmt.Printf("%s %s\n", word, lw)
		}
	}
}

func leet(word string) []string {
	var distWords []string
	matchedBytesMap := make(map[byte]struct{})
	var matchedBytes []byte
	for i := 0; i < len(word); i++ {
		if _, has := leetMap[word[i]]; !has {
			continue
		}
		if _, has := matchedBytesMap[word[i]]; has {
			continue
		}
		matchedBytesMap[word[i]] = struct{}{}
		matchedBytes = append(matchedBytes, word[i])
	}
	fmt.Printf("find %d matches byte in %s\n", len(matchedBytes), word)
	for i := 0; i < len(matchedBytes); i++ {
		depth1 := strings.ReplaceAll(word, string(matchedBytes[i]), leetMap[matchedBytes[i]])
		distWords = append(distWords, depth1)
		for j := i + 1; j < len(matchedBytes) && i != j; j++ {
			distWords = append(distWords, strings.ReplaceAll(depth1, string(matchedBytes[j]), leetMap[matchedBytes[j]]))
		}
	}
	return distWords
}
