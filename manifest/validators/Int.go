package validator

// Int - A signed integer Yaml Manifest Validator Parameter object
//
// This struct allows the consumer to parse the manifest yaml file and to
// verify the manifest file's minimum and maximum threshold values using
// the Verify() method.
//
// A YAML Manifest Validator Parameter object is an implementation of the ParameterInterface type.
type Int struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}
