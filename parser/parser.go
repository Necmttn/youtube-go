package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type XMLTranscript struct {
	XMLName xml.Name  `xml:"transcript"`
	Texts   []XMLText `xml:"text"`
}

type XMLText struct {
	XMLName xml.Name `xml:"text"`
	Context string   `xml:",innerxml"`
	Start   string   `xml:"start,attr"`
	Dur     string   `xml:"dur,attr"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseFile(p string) XMLTranscript {
	// data, err := ioutil.ReadFile(p)
	xmlFile, err := os.Open(p)
	check(err)
	fmt.Println("Successuflly Opened", p)
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	check(err)
	// fmt.Printf("%s", data)
	var altyazi XMLTranscript
	err = xml.Unmarshal(byteValue, &altyazi)
	check(err)
	fmt.Println("Number of Lines: ", len(altyazi.Texts))
	return altyazi
}
