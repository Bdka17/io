package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
	"sigs.k8s.io/yaml"
)

const (
	IO_SCHEMA  = "https://schema.checkpt.in/io_helm_schema.json"
	INPUT_FILE = "aks-manifest/values.yaml"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	jsonInputData := getJsonInputData()

	validateJsonSchema(jsonInputData)

	fmt.Println("Validation Successful")
}

func validateJsonSchema(jsonInputData string) {
	schemaLoader := gojsonschema.NewReferenceLoader(IO_SCHEMA)
	documentLoader := gojsonschema.NewStringLoader(jsonInputData)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	check(err)

	if !result.Valid() {
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		panic("The document is not valid.")
	}
}

func getJsonInputData() (json string) {
	fp, err := filepath.Abs(INPUT_FILE)
	check(err)

	fileBuf, err := os.ReadFile(fp)
	check(err)

	jsonOp, err := yaml.YAMLToJSON(fileBuf)
	check(err)

	return string(jsonOp)
}
