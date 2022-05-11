package main

import (
	"errors"
	"fmt"
	"reflect"
)

type HisPageConfig map[string]interface{}

func (d HisPageConfig) FgShowDocument() (fg int) {
	fg = 1
	if v, ok := d["fg_show_document"]; ok {
		fg, _ = v.(int)
	}

	return
}

func (d HisPageConfig) FgShowPayment() (fg int) {
	fg = 1
	if v, ok := d["fg_show_payment"]; ok {
		fg, _ = v.(int)
	}

	return
}

func GetValueFromISlice(opts []interface{}, parseTo interface{}) error {
	parseToVal := reflect.ValueOf(parseTo)

	if parseToVal.Kind() == reflect.Ptr {
		parseToVal = parseToVal.Elem()
	} else {
		return errors.New("cannot set to unaddressable value")
	}

	for _, opt := range opts {
		optVal := reflect.ValueOf(opt)
		if optVal.Type() == parseToVal.Type() && parseToVal.CanSet() {
			parseToVal.Set(optVal)
		}
	}

	return nil
}

func main() {
	x := make(HisPageConfig)
	x["fg_show_document"] = 1
	fmt.Printf("%#v\n", x.FgShowDocument())
	fmt.Printf("%#v\n", x.FgShowPayment())

	opts := []interface{}{
		"asdfasfd",
		x,
		2342,
	}

	//y := make(HisPageConfig)
	var y HisPageConfig

	GetValueFromISlice(opts, &y)

	fmt.Printf("Y %#v\n", y)
}
