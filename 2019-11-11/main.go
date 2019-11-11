package main

import (
	"fmt"

	//"my_module"
)

type skills interface {
	run()
	speak() string
}
type person struct {
	name string
	age uint
}

func (p person) speak() string{
	return p.name
}

func (person) run(){
	fmt.Println("person.run()")
}

func test(inter skills)  {
	fmt.Println(inter.speak())
}

//method接口，struct 实现了func之后，则该方法变成了struct的私有属性
func (p person) getAge() uint{
	fmt.Println(p.age)
	return p.age
}

func double_return() (string, bool){
	return "string", true
}
func main()  {
	person1 := person{"xiaoming", 15}
	fmt.Printf(person1.name)
	fmt.Printf("main function")

	fmt.Printf(person1.speak())

	test(person1)

	age := person1.getAge()

	fmt.Println(age)

	source := make([]string, 4, 5)
	fmt.Println(len(source), cap(source))
	source = append(source, "first")

	fmt.Println(source, len(source), cap(source))


	//make 用于初始化slice或者ap,
	maps := make(map[string]uint)
	maps["first"] = 1
	maps["second"] = 2

	map2s := make(map[int]string)
	map2s[1] = "first"
	//空的interface代表void类型
	interfaces := make([]interface{}, 10, 10)
	interfaces[0] = maps
	interfaces[1] = map2s

	fmt.Println(interfaces)

	for i, n := range interfaces {

		fmt.Println(i, n)
	}

	for i := 0; i <= 5; i++ {
		if i == 3 || i == 4 {
			fmt.Println(i)
		}

	}

	test_map := make(map[string]int)
	test_map["first"] = 1
	test_map["first"] = 2

	//用于获取value,第二个参数代表是否存在 true/false
	k := test_map["first"]
	k2, v2 := test_map["sss"]
	fmt.Println(k)
	fmt.Println(k2, v2)

	returnStr, _ := double_return()

	fmt.Println(returnStr)
	delete(test_map, "first")
}
/*
总结：
	base datatype:
	支持基本的数据类型，bool, int ,float, string,赋值方式 var name type ,初始化可以使用 var name type = value或者
	name := value，建议使用后者

	func:
	func可以定义method和function两种，当代表method的时候，func (p struct xx) name (param...) return type{...}
	当定义Method的时候，代表该method成为了struct的成员方法。
	func也可以定义为function，当定义为function时 fun name(params...) return type {...}，与普通函数一致
	需要注意，大括号需要同行，可以有多个返回值

	slice:
	slice为动态数据，可以使用make初始化  slices = make([]uint, size, capacity)...类似于char buf[]

	struct:
	struct 为类型的集合和c/c++相似。type my_struct struct{...}

	interface:
	interface 为空接口的集合，当某个struct实现了interface的所有接口，则该struct可以作为interface的参数，例如。func test()

	map:
	一系列的Key:value集合，定义方法为 map1 := map[string]string或者 map2 := make(map[string]string),map的相关方法
	赋值：
	map[key] = value，重复的Key会被后面的value覆盖掉， value, exists := map[key]，exists可选的，表示是否存在
	delete(map, key)用于删除指定key的map,重复删除和删除不存在的key不会报错

*/