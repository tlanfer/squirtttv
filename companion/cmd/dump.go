package main

import (
	"bytes"
	"companion"
	"gopkg.in/yaml.v3"
)

func dump(c companion.Config) string {
	if c.Streamlabs != "" {
		c.Streamlabs = "REDACTED"
	}
	buffer := bytes.Buffer{}
	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(2)
	encoder.Encode(c)
	return buffer.String()
}
