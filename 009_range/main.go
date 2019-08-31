package main
import "fmt"
func main(){
  ids:=[]int{67,48,79,05,96}
  //loop through ids
  for i,id:=range ids{
    fmt.Printf("%d-ID:%d\n",i,id)
  }
  //not using index
  //for i,id:=range ids{
    //fmt.Printf("ID:%d\n",id)
  //}
  //Add ids together
  sum:=0
  for _,id:=range ids{
    sum+=id
  }
  fmt.Println(sum)
  //Range with map
  emails:=map[string]string{"Bob":"bob@gmail.com","shanon":"shanon@gmail.com","ram":"ram00@gmail.com"}
  for k,v:=range emails{
    fmt.Printf("%s: %s\n",k,v)
  }
  for k:=range emails{
    fmt.Println("Name: "+k)
  }


}
