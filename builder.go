package main

import (
	"fmt"
)

type builder interface {
	makeTitle(title string) string
	makeString(str string) string
	makeItems(items []string) string
	close() string
}

type builder2 interface {
	send() string
}

type Director struct {
	builder builder
}

func (self *Director) Construct() string {
	result := self.builder.makeTitle("Title")
	result += self.builder.makeString("String")
	result += self.builder.makeItems([]string{
		"Item1",
		"Item2",
	})
	result += self.builder.close()
	return result
}

type TextBuilder struct {
	joni string
}

func (self TextBuilder) makeTitle(title string) string {
	return "# " + title + "\n"
}

func (self TextBuilder) makeString(str string) string {
	return "## " + str + "\n"
}

func (self *TextBuilder) makeItems(items []string) string {
	var result string
	for _, item := range items {
		result += "- " + item + "\n"
	}
	return result
}

func (self *TextBuilder) close() string {
	return "\n"
}

type infoBuilder struct {}
func(self *infoBuilder) send() string {
	return " ini send ";
}

type animals struct{}
type dogs struct{}


func main() {
	var c builder2
	var j *infoBuilder
	//j = &{}
	c = j//&infoBuilder{} 


	describe(c)
	//var a builder = &TextBuilder{joni: "indo"}
	//d := Director{builder: a}
	d := Director{&TextBuilder{}}
	res := d.Construct()
	fmt.Println(res);
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
