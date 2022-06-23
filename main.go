package main
import ("fmt";"strings";"strconv")

func main() {
    input := "1 9 3 4 -5"
    var x int
    var result string
    
    splitInput:=strings.Split(input," ")
    if len(splitInput)==1{
        result=splitInput[0]
        fmt.Println(result)
        return
    }
    max,_ := strconv.Atoi(splitInput[0])
    min,_ := strconv.Atoi(splitInput[0])
    for i:=0; i<len(splitInput); i++{
       x,_= strconv.Atoi(splitInput[i])
       if max<x{
           max = x 
       }
        if min>x{
           min = x 
       }
    }
    result = strconv.Itoa(max)+" "+strconv.Itoa(min)
    fmt.Println(result)
}
