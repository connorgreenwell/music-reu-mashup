package main

import (
  "encoding/xml"
  "fmt"
  "os"
  "io/ioutil"
)

type Table struct {
  Rows []Row `xml:"row"`
}

type Row struct {
  Columns []Column `xml:"column"`
}

type Column string

func main() {
  xml_file, err := os.Open("nsf.xml")
  if err != nil {
    panic(err)
  }
  defer xml_file.Close()

  xml_data, err := ioutil.ReadAll(xml_file)
  if err != nil {
    panic(err)
  }

  table := Table{}
  err = xml.Unmarshal(xml_data, &table)
  if err != nil {
    panic(err)
  }

  for _, row := range table.Rows {
    fmt.Println(row.Columns[0])
    fmt.Println("\t", row.Columns[3])
    fmt.Println("\t", row.Columns[13])
  }
}
