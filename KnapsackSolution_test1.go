package main

import (
    "fmt"
    "time"
   "runtime"
   "sync"
   "bufio"
    //"log"
    "os"
  "strconv"
    //"knapsackRec"
)


/*
Andie SAMADOULOUGOU 300209487 CSI 2520

Note : 
La fonction Knapsack a ete modifie. Elle declenche des go routines lorsqu'elle doit attendre des resultats pour comparaison 
Cela permet de respecter le paradigme concurrent


*/




/* A brute force recursive implementation of 0-1 Knapsack problem 
modified from: https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10 */

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

// Returns the maximum value that 
// can be put in a knapsack of capacity W 
func KnapSack(W int, wt []int, val []int ) int { 

    // Base Case 
    if (len(wt) == 0 || W == 0) {
        return 0 
    }
    last := len(wt)-1

    // If weight of the nth item is more 
    // than Knapsack capacity W, then 
    // this item cannot be included 
    // in the optimal solution 
    if wt[last] > W { 
    
        return KnapSack(W, wt[:last], val[:last])    

    // Return the maximum of two cases: 
    // (1) nth item included 
    // (2) item not included 
    } else {
        if ( len(wt) == 2){ // s'il reste juste 2 items, utiliser une approche recursive
           return Max(val[last] + KnapSack(W - wt[last], wt[:last], val[:last]), KnapSack(W, wt[:last], val[:last]))
        } else{
            //approche concurrente 
            left := 0
            right := 0
            var wg sync.WaitGroup
            wg.Add(2)
            go func(W int, wt []int, val []int , wg *sync.WaitGroup,value int){
                left = value + KnapSack(W,wt,val)
                wg.Done()
            }(W - wt[last], wt[:last], val[:last],&wg,val[last])
            go func(W int, wt []int, val []int , wg *sync.WaitGroup){
                right =  KnapSack(W,wt,val)
                wg.Done()
            }(W, wt[:last], val[:last],&wg)
           


            wg.Wait()
            max := Max(left,right)
            return max 
        }
       
    }
} 


func main() {

    //obtenir le nom du fichier 
    fmt.Println("Veuiller entrez le nom du fichier d'entree : ")
    var file string
    fmt.Scanln(&file) 

    //variables
    var names []string
    var weights []int
    var values[] int 
    var data [] string // pour garder les infos contenus dans le fichier 
    //var trace[] int //pour garder les  indices des objets ajoutes dans

    //obtenir les informations contenues dans le fichier () // mettre dans une fonction
     f, err := os.Open(file)
    if err != nil {
        fmt.Println(err)
     }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)
    
  
    
    for scanner.Scan() {
        jeton := scanner.Text()
        data = append(data,jeton)
       
    }
    
    //recuperer le nombre d'items 
    N, err := strconv.Atoi(data[0]) //lire le nombre d'items
    //fmt.Println(N)
    W,err := strconv.Atoi(data[len(data)-1])
    //recuperer les noms 
    indexN := 1
    for j := 0 ; j < N ; j++{

        names = append(names,data[indexN])
        indexN += 3
    }
    //recuperer les valeurs 
    indexV := 2
    for j := 0; j< N ; j++{
        t, err := strconv.Atoi(data[indexV])
        values = append(values,t)
        indexV += 3
        if err != nil {
            fmt.Println(err)
        }
    }
    //recuper les poids 
    indexP :=  3
    for j := 0 ; j<N ; j++{
        t, err := strconv.Atoi(data[indexP])
        weights = append(weights,t)
        indexP += 3
        if err != nil {
            fmt.Println(err)
        }
        
    }

   // fmt.Println(names)
   // fmt.Println(values)
   // fmt.Println(weights)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }


    fmt.Println("Number of cores: ",runtime.NumCPU())
    
    // simple example
    //W:= 85//7
    //weights := []int{10, 20,7, 30,3,25}  // {1,2,3,5}
    //values := []int{60, 100,500, 120,100,10}               // {1,6,10,15}
    
    start := time.Now();
    var wg sync.WaitGroup
    wg.Add(2)
    last := len(weights)-1 
    r1 := 0 
    r2 := 0

    //lancer deux go routines respectivement pour l'arbre de gauche et celui de droite
    
    // best solution with the last item
    go func(W int, wt []int, val []int , wg *sync.WaitGroup) {
          //fmt.Println("1")
          r1 = KnapSack(W, wt, val)

          //fmt.Println(r1)
          wg.Done()

    }(W - weights [last], weights [:last], values [:last],&wg) 
    // best solution without the last item
    go func (W int, wt []int, val []int , wg *sync.WaitGroup){
            // fmt.Println("2")
             r2  = KnapSack(W, wt, val)
            //fmt.Println(r2)
            wg.Done()

    }(W, weights, values, &wg) 
  // go KnapSack(W, weights [:last], values [:last], &wg)
   // code here to synchronize and then determine which one is the best solution
       
    toSave := Max(r1,r2)
    wg.Wait()
    if(r1>r2){
        toSave = r1
    } else{
        toSave = r2
    }
    fmt.Println(toSave)
    /*fmt.Println("r1") 
    fmt.Println(r1)
    fmt.Println("r2")
    fmt.Println(r2)*/

    end := time.Now();
    fmt.Printf("Total runtime: %s\n", end.Sub(start))


    //methode de sauvegarde 
    save(file,names,values,weights,toSave)
}

func save(fileName string ,names []string, values []int, weights []int, optimalValue int){
    output := "output_"
    output += fileName
    file, err := os.OpenFile(output, os.O_CREATE|os.O_APPEND, 0600)
    defer file.Close() // on ferme automatiquement le fichier après l'avoir manipulé
    

     _, err = file.WriteString("value optimal     ") // écrire dans le fichier
    if err != nil {
        panic(err)
    }
    
    t := strconv.Itoa(optimalValue)
     _, err = file.WriteString(t) // écrire dans le fichier
    if err != nil {
        panic(err)
    }


}