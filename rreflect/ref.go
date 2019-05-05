package rreflect

import (
	"fmt"
	"reflect"
)

type Job struct {
	Name    string
	Company string
}

type Person struct {
	Name string
	Age  int
	Job  Job
}

func (p *Person) PrintPerson() {
	fmt.Println(p.Name, p.Age)
	p.Job.printJob()
}

func (p Person) Print(abc string) {
	fmt.Println(abc)
}

func (j *Job) printJob() {
	fmt.Println("name is:", j.Name, "com is:", j.Company)
}

func methodOfType() {
	per := Person{"Snow", 37, Job{"King", "Kingdom"}}
	// 获取Type对象
	typPer := reflect.TypeOf(per)
	// rreflect.Person
	fmt.Println(typPer)
	// 获取该接口的具体分类
	// struct
	fmt.Println(typPer.Kind())
	// 该类型在自身包内的类型名
	// Person
	fmt.Println(typPer.Name())
	// 该类型的包路径 明确指定包的import路径
	// goRpc/rreflect
	fmt.Println(typPer.PkgPath())
	// 返回struct类型的字段数（匿名字段算作一个字段），如非结构体类型将panic
	// 3
	fmt.Println(typPer.NumField())
	// 返回struct类型的第i个字段的类型，如非结构体或者i不在[0, NumField())内将会panic
	// {Age  int  16 [1] false}
	fmt.Println(typPer.Field(1))
	// PkgPath is:  Name is: Age Type is: int Index is: [1] Anonymous is: false Offset is: 16 Tag is:
	fmt.Println("PkgPath is: ", typPer.Field(1).PkgPath,
		" Name is: ", typPer.Field(1).Name,
		" Type is: ", typPer.Field(1).Type,
		" Index is: ", typPer.Field(1).Index,
		" Anonymous is: ", typPer.Field(1).Anonymous,
		" Offset is: ", typPer.Field(1).Offset,
		" Tag is: ", typPer.Field(1).Tag)
	// {Name  string  0 [0] false}
	fmt.Println(typPer.FieldByIndex([]int{0}))
	// {Job  rreflect.Job  24 [2] false} true
	fmt.Println(typPer.FieldByName("Job"))
	// 1
	fmt.Println(typPer.NumMethod())
	// {Print  func(rreflect.Person, string) <func(rreflect.Person, string) Value> 0}
	fmt.Println(typPer.Method(0))
	// {Print  func(rreflect.Person, string) <func(rreflect.Person, string) Value> 0} true
	fmt.Println(typPer.MethodByName("Print"))
}

func methodOfValue() {
	per := Person{"hei", 4, Job{"ha", "yes"}}
	valPer := reflect.ValueOf(per)
	fmt.Println(valPer.Kind())
	fmt.Println(valPer.NumMethod())
}

func Ref() {
	methodOfValue()
}
