package lib

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"unicode"
)

var keywords = map[string]int8{
	"break": 1, "default": 1, "func": 1, "interface": 1, "select": 1,
	"case": 1, "defer": 1, "go": 1, "map": 1, "struct": 1,
	"chan": 1, "else": 1, "goto": 1, "package": 1, "switch": 1,
	"const": 1, "fallthrough": 1, "if": 1, "range": 1, "type": 1,
	"continue": 1, "for": 1, "import": 1, "return": 1, "var": 1,
}

var bucket map[string]int

var bucketLocker sync.Mutex

func init() {
	bucket = make(map[string]int)
}

func AddToBucket(word string) {
	defer bucketLocker.Unlock()
	bucketLocker.Lock()
	bucket[strings.ToLower(word)] += 1
}

func GetBucket() map[string]int {
	return bucket
}

func WC(file string) {
	content := ReadFile(file)
	//log.Println(content)
	words := strings.FieldsFunc(content, filter)
	log.Println(words)

	for _, word := range words {
		if !isKeyword(word) {
			AddToBucket(word)
		}
	}
	log.Println(bucket)
}

func isKeyword(word string) bool {
	_, ok := keywords[word]
	return ok
}

func filter(c rune) bool {
	return !unicode.IsLetter(c)
}

func ReadFile(file string) string {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}
	return bytes.NewBuffer(buffer).String()
}

func ReadDir(name string) {
	files, err := ioutil.ReadDir(name)
	if err != nil {
		log.Fatalln(err)
	}
	var fullName string
	for _, file := range files {
		fullName = name + "/" + file.Name()
		//log.Println(fullName, name, file.Name())
		if !file.IsDir() {
			if strings.HasSuffix(fullName, ".go") {
				WC(fullName)
			}
		} else {
			ReadDir(fullName)
		}
	}
}

func WordCount(done <-chan bool) {
	var finished bool
	log.Println("Start to count words...")
	for {
		var repo interface{}
		if finished {
			repo = RepoQueue.Peek()
			if repo == nil {
				break
			}
		}

		select {
		case <-done:
			finished = true
		}
		repo = RepoQueue.Poll()
		if repo != nil {
			path := CLONE_PATH + "/" + repo.(string)
			log.Println("Start count " + path)
			go ReadDir(path)
		}
	}
	log.Println("Finished all word count task....")
}
