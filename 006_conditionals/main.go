package main
import "fmt"
func main(){
  x:=5
  y:=20
  //if else
  if x<y{
    fmt.Printf("%d is less than %d\n",x,y)
  }
   color:="yellow"
  if color=="red"{
      fmt.Println("color is red")
    } else if color=="blue"{
      fmt.Println("color is blue")
  }else{
      fmt.Println("color is not blue or red")
  }
  //else if

  //switch
  switch color{
  case "red":
    fmt.Println("color is red")
  case "yellow":
    fmt.Println("color is yellow")
  case "green":
    fmt.Println("color is green")
  default:
    fmt.Println("color is different")
  }


}
