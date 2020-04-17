package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./input")
	if err != nil {
		fmt.Println("Error while reading input files", err)
		return
	}
	for _, fi := range files {
		name := fi.Name()
		fmt.Println(name)

		f, err := os.Open("./input/" + name)
		if err != nil {
			fmt.Println("Error while reading file", err)
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		var c *Cache
		for scanner.Scan() {
			line := scanner.Text()
			if c == nil {
				split := strings.Split(line, " ")
				bitsPerWord, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println("Error while parsing description", err)
					return
				}
				size, err := strconv.Atoi(split[1])
				if err != nil {
					fmt.Println("Error while parsing description", err)
					return
				}
				bytesPerBlock, err := strconv.Atoi(split[2])
				if err != nil {
					fmt.Println("Error while parsing description", err)
					return
				}
				linesPerBlock, err := strconv.Atoi(split[3])
				if err != nil {
					fmt.Println("Error while parsing description", err)
					return
				}
				c = NewCache(bitsPerWord, size, bytesPerBlock, linesPerBlock)
				continue
			}

			a, err := ParseAddress(line, c)
			if err != nil {
				fmt.Println("Error while parsing address", err)
				return
			}
			if c.access(a) {
				fmt.Print("C")
			} else {
				fmt.Print("M")
			}
		}
		fmt.Println("")
	}
}
