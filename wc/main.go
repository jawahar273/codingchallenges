package main

import (
	"flag"
	"log"
	"os"
  "bufio"
  "strings"
  "fmt"
)

func getLineNumber(file *os.File) int {
  sc := bufio.NewScanner(file)
  result := 0
  for sc.Scan() {
    result += 1
  }
  return result
}

func main() {
 var flagLine = flag.String("l", "", "provide the number of line in the given file") 
  flag.Parse()

  f, err := os.OpenFile("./text.txt", os.O_RDONLY, os.ModeNamedPipe)
  if err != nil {
    log.Fatal("failed to open given file")
    return
  }
  
    var sb strings.Builder
  if flagLine != nil {

    lines := getLineNumber(f)
    sb.WriteString(fmt.Sprint(lines))
    
  }
sb.WriteString(" "+f.Name())
    log.Print(sb.String())
 
}
