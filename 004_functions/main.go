package main
import "fmt"
func greeting(name string)string{
  return "Hello "+name
}
func getsum(a int,b int)int{

  return a+b
}
func main(){
  fmt.Println(greeting("Athira"))
  fmt.Println(getsum(34,25))
}
