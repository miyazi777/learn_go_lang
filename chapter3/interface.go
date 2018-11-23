package main

import (
	. "fmt"
)

// Interfaceを宣言
type Accessor interface {
  GetText() string
  SetText(string)
}

type Document struct {
  text string
}

func (d *Document) GetText() string {
  return d.text
}

func (d *Document) SetText(text string) {
  d.text = text
}

func main() {
  var doc *Document = &Document{}
  doc.SetText("document")
  Println(doc.GetText())

  var acsr Accessor = &Document{}
  acsr.SetText("accessor")
  Println(acsr.GetText())
}
