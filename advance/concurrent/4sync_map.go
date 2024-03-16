package concurrent

import (
	"fmt"
	"sync"
)

func rangeMap(_map *sync.Map) {
	_map.Range(func(key, value any) bool {
		fmt.Printf("key is %v, val is %v\n", key, value)
		return true
	})
	fmt.Println("------")
}
func main() {
	var m sync.Map

	m.Store("name", "zhangshan")
	m.Store("age", 18)

	age, _ := m.Load("age")
	fmt.Println(age)
	fmt.Println(age.(int))

	rangeMap(&m)

	m.Delete("age")

	rangeMap(&m)

	age, _ = m.Load("age")
	fmt.Println(age)

	//m.Delete("name")  // 删除后，LoadOrStore loaded 返回 false

	name, loaded := m.LoadOrStore("name", "aa")
	fmt.Println(name, loaded)

}
