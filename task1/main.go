package main
import "fmt"

func main() {
    arr := []int{4, 1, 4, -4, 6, 4, 3, 8, 8}
    a1:=make([]int, len(arr), len(arr))
    var k int
    var result []int
    
    //делаем копию масива(она понадобиться в конце))))))
    for i:=0; i<len(arr); i++{//
        a1[i]=arr[i]
    }
    
    ch:=(-11)
    
    //узнаем какого єлемента нет в массиве 
    for j:=-11; j<10000; j++{
        for i:=0; i<len(arr); i++{//
            if arr[i]==ch{
                ch++
                break
            }
        }
    }
    
    //заменяем повторяющ єлементы на несуществующий 
    for i:=0; i<len(arr)-1; i++{
        for j:=i+1; j<len(arr); j++{
            if arr[i]==arr[j]{
                arr[j]=ch
            }
        }
    }
    
    //узнаем количество повторяющихся єлементов 
    for i:=0; i<len(arr); i++{
        if arr[i]==ch{
            k++
        }
    }
    
    //создаем новый масив без повторяющихся єлементов
    aa:=make([]int, len(arr)-k, len(arr)-k)
    jj:=0
    for i:=0; i<len(arr); i++{
        if arr[i]==ch {
            continue
        }
        aa[jj]=arr[i]
        jj++
    }
    result=aa
    fmt.Println(result)
}
