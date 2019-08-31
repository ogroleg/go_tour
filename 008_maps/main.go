package main
import "fmt"


func main(){
  //define maps
  //emails:=make(map[string]string)
  //assign kv
  //emails["bob"]="bob@gmail.com"
  //emails["sharon"]="sharon@gmail.com"
  //emails["ram"]="ram@gmail.com"
  //declare map0 and assign kv
  emails:=map[string]string{"Bob":"bob@gmail.com","shanon":"shanon@gmail.com","ram":"ram00@gmail.com"}
  fmt.Println(emails)
  fmt.Println(len(emails))
  fmt.Println(emails["ram"])
  delete(emails,"Bob")
  //delete
  emails["mike"]="mike34@gmail.com"
  fmt.Println(emails)


}
