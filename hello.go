package main

import (
	"fmt"
	"log"
	// "github.com/wizardkeven/stringutil"
	// "math"
	"bufio"
	"os"
	"strings"
	"strconv"

)

func main(){			
	loop,_:= getInput(0,0)
	// fmt.Printf("Loop times: %d\n", loop)
	callForSub(loop)

}

func getInput(tt,lp int) (int, []int){
	var inp int
	var ina []int
	if tt == 0 {
		_, err:= fmt.Scanf("%d", &inp)
		if err != nil{
			log.Fatal(err)
			// fmt.Printf("Loop times: %d\n", loop)
		}

	}else if tt == 1{
		// var text string
		reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		// fmt.Scanln(&text)
		// fmt.Println("getInput para:",text,"\n")
		oo := strings.Fields(text)
		// ll := len(oo)
		// fmt.Sscanf(text, "%d", ina) 
		if len(oo)!=lp {
			// fmt.Printf("Entered numbers amount not correct!\n")
			return 0,nil
		}
		ina = getInArray(oo)
		// fmt.Println(ina)
	}
	return inp,ina

}
func getInArray(str []string) []int{
	if len(str)<=0 {
		return nil
	}
	if len(str)==1 {
		var result [2]int
		if tem, err := strconv.Atoi(str[0]); err == nil {
			result[0] = tem
			return result[0:1]
		}else{
			// fmt.Printf("Number format not correct: %s\n", str[0])
			// return nil
			log.Fatal(err)
		}
	}else{
		return append(getInArray(str[0:len(str)/2]),getInArray(str[len(str)/2:])...)
	}
	return nil
}
func getPow(ar []int) int{
	if len(ar)<=0 {
		return 0
	}
	if len(ar) == 1{
		if ar[0]>0 {
			return ar[0]*ar[0]
		}else{
			return 0
		}
	}else{
		return getPow(ar[0:len(ar)/2])+getPow(ar[len(ar)/2:])
	}
}

func callForSub(num int){
	if num >0 {
		//get input and calculate
		var subLoop int
		subLoop,_= getInput(0,0)
		// fmt.Printf("Sub loop times: %d\n", subLoop)
		if subLoop >0 {
			var arr []int
			_,arr=getInput(1,subLoop)
			// fmt.Printf("Sub arrays\n")
			// fmt.Println(arr)
			if arr == nil {
				return
			}
			fmt.Println(getPow(arr))

		}else{
			fmt.Printf("Cannot enter %d numbers!\n", subLoop)
		}

		// fmt.Print(num,"\n")
		num--
		callForSub(num)
	}
}
