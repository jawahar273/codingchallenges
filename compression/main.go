package main

import (
	"bufio"
	"log"
	"os"
)

func RuneOccurance(occurance map[rune]int, val string)  {
  for _, rune := range  val {
    occurance[rune]++
  }
}


func main() {
  file, err := os.OpenFile("./compression/victor-hugo.txt", os.O_RDONLY, 0644) 
  if err != nil {
   log.Fatal("error on open file", err) 
  }
  defer file.Close()
      scanner := bufio.NewScanner(file)
     occurance := make(map[rune]int)

    for scanner.Scan() {
        // do something with a line
        // fmt.Printf("line: %s\n", scanner.Text())
        RuneOccurance(occurance,  scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    log.Println("--", occurance)
}
