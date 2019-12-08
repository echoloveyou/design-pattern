package factory_pattern

import (
	"fmt"
	"strings"
)

type PizzaType int8

const (
	_                  PizzaType = iota
	CheesePizzaType              //奶酪披萨
	ClamPizzaType                //蛤蜊披萨
	PepperoniPizzaType           //香肠披萨
)

type IPizza interface {
	Prepare() //制作披萨的准备工作
	Bake()    //烘烤
	Cut()     //切片
	Box()     //打包
}

type Pizza struct {
	Name     string
	Dough    string   //面团
	Sauce    string   //酱
	Toppings []string //配菜
}

//纽约奶酪披萨
type NyCheesePizza struct {
	Pizza
}

//纽约火腿披萨
type NyPepperoniPizza struct {
	Pizza
}

//中国蛤蜊披萨
type ChinaClamPizza struct {
	Pizza
}

//中国奶酪披萨
type ChinaCheesePizza struct {
	Pizza
}

//中国火腿披萨
type ChinaPepperoniPizza struct {
	Pizza
}

//纽约蛤蜊披萨
type NyClamPizza struct {
	Pizza
}

func NewNyCheesePizza() NyCheesePizza {
	return NyCheesePizza{
		Pizza{
			Name:     "奶酪披萨",
			Dough:    "白面面团",
			Sauce:    "沙拉酱",
			Toppings: []string{"奶酪"},
		},
	}
}

func NewNyPepperoniPizza() NyPepperoniPizza {
	return NyPepperoniPizza{Pizza{
		Name:     "纽约火腿披萨",
		Dough:    "白面面团",
		Sauce:    "番茄酱",
		Toppings: []string{"火腿"},
	}}
}

func NewNyClamPizza() NyClamPizza {
	return NyClamPizza{Pizza{
		Name:     "纽约蛤蜊披萨",
		Dough:    "白面面团",
		Sauce:    "番茄酱",
		Toppings: []string{"蛤蜊", "生菜"},
	}}
}

func NewChinaCheesePizza() NyCheesePizza {
	return NyCheesePizza{
		Pizza{
			Name:     "中国奶酪披萨",
			Dough:    "白面面团",
			Sauce:    "沙拉酱",
			Toppings: []string{"奶酪"},
		},
	}
}

func NewChinaPepperoniPizza() NyPepperoniPizza {
	return NyPepperoniPizza{Pizza{
		Name:     "中国火腿披萨",
		Dough:    "白面面团",
		Sauce:    "番茄酱",
		Toppings: []string{"火腿"},
	}}
}

func NewChinaClamPizza() NyClamPizza {
	return NyClamPizza{Pizza{
		Name:     "中国蛤蜊披萨",
		Dough:    "白面面团",
		Sauce:    "番茄酱",
		Toppings: []string{"蛤蜊", "生菜"},
	}}
}

func (p Pizza) Prepare() {
	fmt.Printf("正在准备制作：%s口味的披萨\n", p.Name)
	fmt.Printf("加入面团:%s\n", p.Dough)
	fmt.Printf("加入酱:%s\n", p.Sauce)
	fmt.Printf("配菜加入:%s\n", strings.Join(p.Toppings, ","))
}

func (p Pizza) Bake() {
	fmt.Println("拷贝30分钟")
}

func (p Pizza) Cut() {
	fmt.Println("披萨切片")
}

func (p Pizza) Box() {
	fmt.Println("披萨打包")
}
