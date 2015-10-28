/*
Package for creating and manipulating additive manufacturing files.
*/
package amf

import "encoding/xml"

// Scene is the top level AMF structure containing objects, materials, etc.
type Scene struct {
	XMLName  xml.Name `xml:"amf"`
	Objects  []*Object
	Material []*Material
	Texture  []*Texture
	Metadata []*Metadata
}

type Object struct {
	ID   string `xml:"id,attr"`
	Mesh []*Mesh
}

type Mesh struct {
	Vertices []*Vertex
	Volumes  []*Volume
}

type Volume struct {
	Triangles []*Triangle
}

type Triangle struct {
	V1 int
	V2 int
	V3 int
}

//type Triangle [3]int
//
//func (t *Triangle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
//	return nil
//}

type Vertex struct {
	X float64 `xml:"x"`
	Y float64 `xml:"y"`
	Z float64 `xml:"z"`
}

type Material struct {
	ID string `xml:"id,attr"`
}

type Texture struct {
}

type Metadata struct {
	Type  string
	Value string
}
