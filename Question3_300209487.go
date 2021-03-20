package main
import "fmt"
import "math"
import "time"
import "math/rand"
import "sync"


/*
L'utilisation des 5 go routines va nous permettre de reduire les taches. 
Nous devons generer 1000 nombres premiers selon l'énoncé. Dans le copde deja fourni , 10 000  nom uniques sont generes (la fonction Unicite que j'ai ajouté permet de le vérifier).
Nous allons alors  confier a chacune des 5 go routines la generation de 200 nombres premniers sur  des intervalles definis.
*/

// returns true if number is prime
func isPrime(v int64) bool {
           sq:= int64(math.Sqrt(float64(v)))+1
           var i int64
           for i=2; i<sq; i++ {
                  if v%i == 0 {
                     return false
                   }
           }
            return true
}


// get a random prime number between 1 and maxP
func getPrime(maxP int64) int64 {
         var i int64
         for i=0; i<maxP; i++ {

             n:= rand.Int63n(maxP)
             if isPrime(n) {
                     return n
             }
         }
         return 1 // just in case
}

func Unicite( tab [] int64) bool {
      for i :=0 ; i < len(tab) ; i++{

              for j := i + 1 ; j< len(tab) ; j++ {
                    if ( tab[i] == tab[j]) {
                           return false          
                     } else {
                    
                     }
              }
       }
       return true
}

//go routines generant les nombres premiers sur un intervalle particulier 

func fill(debut int  , fin int , maxP int64 ) [] int64 {
          var tab []int64
          i := 0 
         for i < 200 {
              p:= getPrime(maxP) // add a new prime
              //tester la valeur pour savoir si elle est dans le bon intervalle
              if( int(p) > debut) {
                 tab = append(tab,p)
                 i++
              } else {
              }
             
         }
         
         
        // wg.Done()
         return tab


} 


func main() {
        //tabs := []int64{9,35,27,56,88,80,88}
       var primes []int64 // slice of prime numbers
        const maxPrime int64 = 10000000 // max value for primes
        start:= time.Now()
        var wg sync.WaitGroup
        wg.Add(5)
       for index := 0 ; index < 5 ; index++{
            
            start_i  := index *250000
            end_i := start_i + 250000
            go func(debut int, fin int , maxP int64){
               defer wg.Done()
               temp := fill(debut,fin,maxP)
               for k := 0 ; k < 200 ; k++ {
                        primes = append(primes,temp[k])
    
               }
               
                


             }(start_i, end_i, maxPrime)      


       }
        wg.Wait()
        end:= time.Now()
      fmt.Println("nombre de nombres premiers : ")
      fmt.Println(len(primes))
      fmt.Println(primes)

        fmt.Println("End of program.",end.Sub(start))
}