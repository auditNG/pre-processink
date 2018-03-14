package transform

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"os"
)

var spaceChar = byte(' ')
var transformConfigPath = string("./transform/transform_config.json")

func NewTransform() Transform {
	return Transform{}
}

type Transform struct {
}

func (t Transform) Process(input string, outputFile *os.File) error {

	jsonparser.ArrayEach([]byte(input),
		func(actVal []byte, _ jsonparser.ValueType, _ int, err error) {
			timestamp, err := jsonparser.GetString(actVal, "_source", "@timestamp")
			if err != nil {
				fmt.Println("JSON parsing error: ", err)
				return
			}
			fmt.Println("timestamp: " + timestamp)

			message, err := jsonparser.GetString(actVal, "_source", "message")
			if err != nil {
				fmt.Println("JSON parsing error: ", err)
				return
			}
			t.processMessage(message, outputFile)
			if nil != err {
				return
			}

		}, "hits", "hits")

	return nil
}

func (t Transform) processMessage(message string, outputFile *os.File) error {

	fimTransformer := NewFIMTransformer()

	config, err := ioutil.ReadFile(transformConfigPath)
	if err != nil {
		fmt.Println("Error reading transform config")
		return err
	}

	fimTransformer.Process(message, string(config), outputFile)

	return nil
}
