package main


import "fmt"
import "sync"
import "time"


func process(v int) int {
 time.Sleep(1500*time.Millisecond) // simulate compute time
 return 2*v
}


func main() {
            var wg sync.WaitGroup
            tab := []int{9,35,27,56,88,80}
            for i:= 0 ; i< len(tab) ;i++ {
                 wg.Add(1)
                 go func(index int) {
                        defer wg.Done()
                        fmt.Println(process(tab[index]))
                 }(i)
            }
            wg.Wait()
}