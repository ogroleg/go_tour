package main
import "fmt"
func main(){
  //Arrays
  //var fruit [2]string
  //Assigning values
  ///fruit[0]="Apple"
  //fruit[1]="Mangoes"
  //fruit :=[2]string{"Apple","Orange"}
  //fmt.Println(fruit)
  //fmt.Println(fruit[1])
  fruitSlice:=[]string{"Apple","Mango","Orange"}
  fmt.Println(len(fruitSlice))
  fmt.Println(fruitSlice[0:2])
}
