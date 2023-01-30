package main
 
import (
    "fmt"
)
 
func main() {
    byteArray := []byte{'G', 'O', 'L', 'A', 'N', 'G'}
    str1 := string(byteArray)
    fmt.Println("String =",str1)
    fmt.Printf("String = %s\n",byteArray)
}
