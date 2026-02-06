package main

import (
	"fmt"
)

func main() {

}

// 任务1：数组倍增器
func task1() {
	arr := [5]int{10, 20, 30, 40, 50} // 只要带容量大小的都是数组，因为数组的大小算是它的一个特性
	doubleArrayByVal(arr)
	fmt.Println("1.错误演示，按值传递的方式：")
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
	fmt.Println("2.正确演示，按指针传递的方式：")
	doubleArrayByPtr(&arr)
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
}

// 错误示范：值传递
func doubleArrayByVal(arr [5]int) {
	for _, v := range arr {
		v *= 2
	}
}

// 正确示范：指针传递
func doubleArrayByPtr(arr *[5]int) {
	// 要求：在函数内部，使用指针解引用或直接索引的方式，将原数组每个元素乘以 2。
	for i := range arr {
		arr[i] *= 2
	}
}

/*
思考题 (写在注释里)：
- 在 `doubleArrayByPtr` 中，访问元素时写 `(*arr)[i]` 和 `arr[i]` 有区别吗？Go 编译器在这里做了什么？
- 答：解引用遍历的本质是 (*arr)[i]，但 Go 允许简化成 arr[i]，Go语言会自动解引用。
*/

// 任务2：动态数据流处理
func task2() {
	// 1.切片删除 (Delete)：
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //没有大小，即为切片，切片是引用类型传递
	data = append(data[:4], 6, 7, 8, 9)
	fmt.Println("删除指定数字5后的切片：")
	for i := range data {
		fmt.Printf("%d ", data[i])
	}

	// 2.切片备份
	backup := make([]int, len(data), cap(data))
	copy(data, backup)
	for i := range data {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println("\n判断数据是否备份正确:")
	for i := range backup {
		fmt.Printf("%d ", data[i])
	}

	// 3.扩容观察
	fmt.Println("扩容观察")
	for i := 0; i < 5; i++ {
		data = append(data, 100)
		fmt.Println("此时的len:%d,cap:%d", len(data), cap(data))
	}
}

// 任务3：即时计票系统
func task3() {
	// 包含重复的名字的投票箱
	votes := []string{"Alice", "Bob", "Alice", "Charlie", "Alice", "Bob", "Charlie", "David", "Alice"}
	counter := make(map[string]int)
	for _, v := range votes {
		_, ok := counter[v] //检查key是否存在：value是对应的值，ok是bool
		if !ok {
			counter[v] = 1
		} else {
			counter[v] += 1
		}
	}
	var (
		max_num  int
		max_name string
	)
	for name, count := range counter {
		fmt.Printf("Candidate [%s]: [Count] %d votes\n", name, count)
		if count > max_num {
			max_num = count
			max_name = name
		}
	}
	fmt.Println("票数最多的候选人为%s,票数为%d", max_name, max_num)
}

/* 综合实战项目：简易库存管理系统,这是一个模拟电商后台的小程序。你需要结合结构体（定义商品）、Map（存储数据）、指针（修改状态）和 函数。 */
type Product struct {
	id    int
	name  string
	price float64
	stock int // 库存数量
}

func project() {
	/*note:
	 *只是声明了一个 nil map，直接赋值会panic，必须用 m := make(...)或者字面量初始化。
	 *var inventory map[int]*Product
	 *inventory = make(map[int]*Product)
	 *字面量初始化（直接写值）
	 *m1 := map[string]int{"Alice": 1, "Bob": 2}
	 *非字面量初始化（用make函数）
	 *m2 := make(map[string]int)
	 */

	iPhone := &Product{101, "iPhone", 5000.00, 20}
	Laptop := &Product{102, "Laptop", 20.00, 5}
	Bag := &Product{103, "Bag", 65.00, 2}
	Basketball := &Product{104, "Basketball", 90.00, 1}
	inventory := map[int]*Product{
		101: iPhone,
		102: Laptop,
		103: Bag,
		104: Basketball,
	}
	showInventory(inventory)
	addProduct(inventory, 105, "Cap", 18.00, 2)
	buyProduct(inventory, 105, 2)
	buyProduct(inventory, 105, 1)
	addProduct(inventory, 105, "Cap", 18.00, 2) //再添加两个库存
	showInventory(inventory)
}

// 模块功能实现
func showInventory(m map[int]*Product) {
	fmt.Println("=== 商品库存清单 ===")
	for _, p := range m {
		fmt.Printf("商品ID: %d\n", p.id)
		fmt.Printf("名称:   %s\n", p.name)
		fmt.Printf("价格:   %.2f\n", p.price)
		fmt.Printf("库存:   %d\n", p.stock)
		fmt.Println("-------------------")
	}
}

func addProduct(m map[int]*Product, id int, name string, price float64, stock int) {
	product, ok := m[id]
	if ok {
		fmt.Println("商品已存在，正在为该商品添加库存")
		product.stock += stock
	} else {
		fmt.Println("商品不存在，正在创建一个新商品")
		newproduct := &Product{id, name, price, stock}
		m[id] = newproduct
	}
}

func buyProduct(m map[int]*Product, id int, count int) {
	product, ok := m[id]
	if ok {
		if product.stock < count {
			fmt.Println("库存不足")
		}
		product.stock -= count
		fmt.Println("购买成功，共购买%d件商品，花费 %f 元\n 该商品还剩余%d件", count, float64(count)*product.price, product.stock)
	} else {
		fmt.Println("商品不存在")
	}
}
