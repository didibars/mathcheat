package main

import (
    "bufio"
    "fmt"
    "log"    
//    "net/url"
    "os"
    "strings"
)

// load htmlFile, read line by line. 
// parse the value of the given attribute
// prepend baseUrl
// replace spaces with + (url encode)
// print the line to the new file
// all other lines - print as is.
func convertLinksToS3(baseUrl, htmlFileIn, attributeName string) {
    f, err := os.Open(htmlFileIn)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
    	fExpectValue := false	
	originalLink := ""
	newLink := ""
	// tokenize by space, then by =, strip quotes.
	tokens := strings.Split(line, "\"")
	//every other odd one is a value
	// every even one is an attribute name plus some other chars
	for _, token := range tokens {
	    if fExpectValue {
	       fExpectValue = false
	       originalLink = token
	       newLink = fmt.Sprintf("%s%s", baseUrl, strings.Replace(token, " ", "+", -1))
	       continue
	    }
	    key := strings.TrimSpace(token)
	    if len(key) < 2 {
	       continue
	    }
	    key = key[:len(key)-1]
	    key = strings.TrimSpace(key)
	    if key == attributeName {
	       // next one will be value
	       fExpectValue = true
	    }
	}
	if len(newLink) > 0 {
	   fmt.Println(strings.Replace(line, originalLink, newLink, 1))
	} else {
	  fmt.Println(line)
	}
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

