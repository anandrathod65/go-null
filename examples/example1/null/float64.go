package null

import (
	"encoding/json"
	"log"
	"runtime/debug"
)

//Auto-generated code; DONT EDIT THIS CODE

type Float64 struct {
	val   float64
	valid bool
}

func NewFloat64(val float64) Float64 {
	return Float64{val: val, valid: true}
}

func (t *Float64) Set(val float64) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *Float64) Get() float64 {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:Float64!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *Float64) GetPtr() *float64 {
	return &t.val
}

func (t *Float64) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *Float64) SetSafe(val float64) {
	t.val = val
	t.valid = true
}

func (t *Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
