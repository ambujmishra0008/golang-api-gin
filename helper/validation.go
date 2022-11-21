package helper

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateRequest(reqData []byte) (bool, error) {

	pwd, err := os.Getwd()
	if err != nil {
		return false, err
	}
	schemaLoader := gojsonschema.NewReferenceLoader("file:///" + pwd + "/validation.json")

	documentLoader := gojsonschema.NewBytesLoader(reqData)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false, err
	}
	// Concatenate errors slice elements in a string
	if !result.Valid() {
		errStr := "JSON is not valid. See errors: "
		for _, desc := range result.Errors() {
			errStr = errStr + fmt.Sprintf("%s, ", desc)
		}
		errStr = strings.TrimRight(errStr, ", ")
		return false, errors.New(errStr)
	}

	return true, nil
}
