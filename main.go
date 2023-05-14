package main

/** 
learn_go_module hanya alias
nama package di module harus pake _ ga blh - dan nama package engga harus sama dengan nama dir nya itu sendiri

kalo ada major changes yg merubah fitur (func) lama, hrs ubah nama module cth: .../v2
*/

import ( 
	"fmt"
	learn_go_module "github.com/Snow-00/learn-go-module/v2"
)

func main() {
	fmt.Println(learn_go_module.SayHello("test"))
}