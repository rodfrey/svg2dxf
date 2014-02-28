package svg2dxf

import (
    "io/ioutil"
    "fmt"
    "encoding/xml"
    )

type SVG struct {

}

func ParseSVGFile(filename string) (*SVG, error) {
    xmlData, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return nil, err
    }

    var svg SVG
    xml.Unmarshal(xmlData, &svg)
   return &svg, nil
}

