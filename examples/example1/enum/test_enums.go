package enum

//go:generate go-null -package=github.com/rimpo/go-null/examples/example1 -output=..

type TypeMapStruct string

type TypeTest TypeMapStruct

type TypeTestMaster struct {
	Field1 string
	Field2 string
	Field3 string
	Field4 string
}

var MapTypeEducationIDToText = map[string]TypeTestMaster{
	"1": TypeTestMaster{Field1: "", Field2: "", Field3: "", Field4: ""},
	"2": TypeTestMaster{Field1: "", Field2: "", Field3: "", Field4: ""},
}
