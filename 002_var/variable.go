//Datatypes:
//string
//bool
//int int8 int16 int32 int64
//uint uint8 uint32 uint64
//byte-alias for uint8
//rune-alias for int32
//float32 float64
//complex64 complex128

//Using var
package main
  import "fmt"
  //var name="Ann"
  var size float32=3.8
  func main(){
  
  name,email := "Ann","ann2000@gmail.com"//cannot declare outside the function like this.In case wanna describe then use var name="Ann"
  var age=10
  var isCool=true
  fmt.Println(name,age,size,email)
  //%T defines datatype
  //const keyword means cannot redefine
  fmt.Printf("%T\n",name)
  fmt.Printf("%T\n",age)
  fmt.Printf("%T\n",isCool)
  fmt.Printf("%T\n",size)
  fmt.Printf("%T\n",email)

}
