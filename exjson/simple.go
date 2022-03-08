package exjson

import (
	"encoding/json"
	"io"
	"os"
)

func ReadList_StringSlice(filename string) ([]string, error) {
	var l []string
	err := readJSON(filename, &l)
	return l, err
}

func ReadDict_StringBoolMap(filename string) (map[string]bool, error) {
	m := map[string]bool{}
	err := readJSON(filename, &m)
	return m, err
}

func Read(filename string, ptrToValue interface{}) error {
	return readJSON(filename, ptrToValue)
}

func readJSON(filename string, v interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	d := json.NewDecoder(f)
	return d.Decode(v)
}

func Save(d interface{}, saveTo string) error {
	f, err := os.OpenFile(saveTo, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	e := json.NewEncoder(f)
	return e.Encode(d)
}

func SaveToFile(d interface{}, f io.Writer) error {
	e := json.NewEncoder(f)
	return e.Encode(d)
}
