package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Stu struct {
	Name    string `xml:"name"`
	Address string `xml:"addr,omitempty"`
	Secret  string `xml:"-"`
	Father  string `xml:"parent>father"`
	Mother  string `xml:"parent>mother"`
	// Par     Parent `xml:"parent"`
	Note string `xml:"note,attr"`
}

// type Parent struct {
// 	Father string `xml:"father"`
// 	Mother string `xml:"mother"`
// }

func main() {
	stu := &Stu{
		Name:   "myName",
		Secret: "mySecret",
		Father: "myFather",
		Mother: "myMother",
		Note:   "MyNote",
	}

	// data, err := xml.Marshal(stu)
	data, err := xml.MarshalIndent(stu, "", "	")
	checkError(err)

	err = ioutil.WriteFile("stu.xml", data, 0644)
	checkError(err)

	content, err := ioutil.ReadFile("stu.xml")
	checkError(err)

	stuNew := &Stu{}
	err = xml.Unmarshal(content, stuNew)
	checkError(err)

	fmt.Printf("content is : %#v\n", stuNew)

	fmt.Println(">>> stuNew.Name: ", stuNew.Name)
}
