//Adapter
//适配器模式
//适用场景：
//		将一个类的接口转换成用户所期待的。
//		适配器使得原本因接口不兼容而不能在一起工作的类工作在一起
//
//《大话设计模式》-篮球翻译适配器

package Adapter

import (
	"fmt"
)

type Player interface {
	Attack()
	Defense()
}

type Forwards struct {
	name string
}

func (f *Forwards) Attack() {
	fmt.Println("前锋 " + f.name + " 进攻")
}

func (f *Forwards) Defense() {
	fmt.Println("前锋 " + f.name + " 防守")
}

func NewForwards(name string) Player {
	return &Forwards{name}
}

type Guards struct {
	name string
}

func (g *Guards) Attack() {
	fmt.Println("后卫 " + g.name + " 进攻")
}

func (g *Guards) Defense() {
	fmt.Println("后卫 " + g.name + " 防守")
}

func NewGuards(name string) Player {
	return &Guards{name}
}

type Center struct {
	name string
}

func (c *Center) Attack() {
	fmt.Println("中锋 " + c.name + "进攻")
}

func (c *Center) Defense() {
	fmt.Println("中锋 " + c.name + " 防守")
}

func NewCenter(name string) Player {
	return &Center{name}
}

type ForeignCenter struct {
	name string
}

func (fc *ForeignCenter) jingong() {
	fmt.Println("外籍中锋 " + fc.name + " 进攻")
}

func (fc *ForeignCenter) fangshou() {
	fmt.Println("外籍中锋 " + fc.name + " 防守")
}

//Translator 适配器，将ForeignCenter接口转换成用户所期待的
type Translator struct {
	fc ForeignCenter
}

func (t *Translator) Attack() {
	t.fc.jingong()
}

func (t *Translator) Defense() {
	t.fc.fangshou()
}

func NewTranslator(name string) Player {
	return &Translator{ForeignCenter{name}}
}
