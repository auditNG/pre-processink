package transform

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"os"

)

var spaceChar = byte('\n')
var transformConfigPath = string("./transform/transform_config.json")

func NewTransform() Transform {
	return Transform{}
}

type Transform struct {
}

func (t Transform) Process(input string, outputFile *os.File) error {

var message=string("")
	jsonparser.ArrayEach([]byte(input),
		func(actVal []byte, _ jsonparser.ValueType, _ int, err error) {

			jsonparser.ArrayEach(actVal, func(value []byte, dataType jsonparser.ValueType, offset int,err error)  {
				//fmt.Println(message[i])
				jsonparser.ObjectEach(value,func(key []byte, pair []byte, dataType jsonparser.ValueType, offset int) error{
					message=message+string(key)+"="+string(pair)+"\n"

					return nil
				})
			


					}, "_source", "osquery_distributed_query_result","result")
					fmt.Println(message)

		}, "hits","hits")
	t.processMessage(message,outputFile)
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
