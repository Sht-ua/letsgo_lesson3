package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Worker struct {
	Person
	Exp int
}
type Driver struct{
	Worker
	Category string
}
type Doctor struct {
	Worker
	Stack string
}
type Policeman struct {
	Worker
	Rank string
}
type Teacher struct {
	Worker
	Degree string
}
type Firefighter struct {
	Worker
	Rank2 string
}
var p1 = Person{"Олег",25}
var p2 = Person{"Игорь",37}
var p3 = Person{"Сергей",24}
var p4 = Person{"Жанна",19}
var p5 = Person{"Ольга",29}
var w1 = Worker{Exp: 5}
var w2 = Worker{Exp: 7}
var w3 = Worker{Exp: 2}
var w4 = Worker{Exp: 4}
var w5 = Worker{Exp: 11}

//var ID1 = map[int]Person{1:{"Вася",37},2 :{"Коля", 42}, 3 :{"Олеся",22}}
//var xs = map[int]Worker{21:{p1, 10}}
//var d1 = map[int]Worker{45:{p2, 2}}
var driverID = map[int]Driver{500:{w1,"Водитель автобуса"}}
var doctorID = map[int]Doctor{600:{w2,"Кардиолог"}}
var policemanID = map[int]Policeman{700:{w3,"Капитан"}}
var teacherID = map[int]Teacher{800:{w4,"Класный руководитель"}}
var firefighterID = map[int]Firefighter{900:{w5,"Дежурный"}}
func main(){
		fmt.Println(p1,driverID)
	    fmt.Println(p2,doctorID)
		fmt.Println(p3,policemanID)
		fmt.Println(p4,teacherID)
		fmt.Println(p5,firefighterID)
}