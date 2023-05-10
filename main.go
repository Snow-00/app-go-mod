package main

/** 
learn_go_module hanya alias
nama package di module harus pake _ ga blh - dan nama package engga harus sama dengan nama dir nya itu sendiri
*/
import ( 
	"fmt"
	learn_go_module "github.com/Snow-00/learn-go-module"
)

func main() {
	fmt.Println(learn_go_module.SayHello())
}