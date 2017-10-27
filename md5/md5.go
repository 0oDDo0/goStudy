package main

import (
	"os"
	"crypto/md5"
	"io"
	"fmt"
)

func main()  {
	testFile := "D:/baiDuYun/Go/goStudy/test.txt"
	file, inerr := os.Open(testFile)
	fmt.Printf("1")
	if inerr == nil {
		md5h := md5.New()
		fmt.Printf("2")
		io.Copy(md5h, file)
		fmt.Printf("md5 = %x", md5h.Sum([]byte(""))) //md5
	}
	defer file.Close()
}
