package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func readFile(path string){
	file, err := os.Open(path)
	if err != nil{
		fmt.Println(path, "not exist")
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		// 包括最后的截止符
		line, err := br.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		fmt.Println("read line", line)
	}
}
func writeFile(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()
	f.WriteString("line 0 \n")
	w := bufio.NewWriter(f)
	strs := []string{"1", "2"}
	for index, str := range strs {
		lineStr := fmt.Sprintf("index = %d, val = %s\n", index, str)
		fmt.Fprintf(w, lineStr)
	}
	return w.Flush()
}

func main() {
	writeFile("test.txt")
	readFile("test.txt")
}
