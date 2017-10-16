package import_test

import "fmt"

func init()  {
	fmt.Println("use init")
}
func PrintUse()  {
	PrintConfig()
	fmt.Println(fmt.Sprintf("use.go,    a = %d b = %d c = %d d = %d", a, b, c, d))
}