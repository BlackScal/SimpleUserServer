package configs

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v2"
)

func Parse(filename string, in interface{}) error {
	yamlContent, err := ioutil.ReadFile(filename)
	if err != nil {
		msg := fmt.Sprintf("ReadFile %s failed", filename)
		return errors.WithMessage(err, msg)
	}
	err = yaml.UnmarshalStrict(yamlContent, in)
	if err != nil {
		msg := fmt.Sprintf("UnmarshalStrict %s failed", filename)
		return errors.WithMessage(err, msg)
	}

	return nil
}
