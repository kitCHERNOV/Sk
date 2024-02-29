package main
 
import (
    "fmt"
    "log"
)
 
func main() {
    var n string // Think that type string better use for other date
    fmt.Print("Введите данные: ")
    _, err := fmt.Scan(&n)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Вы ввели данные: %v\n", n)
}