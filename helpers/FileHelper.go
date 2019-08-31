package helpers

import "io/ioutil"

//Load the file
func LoadFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
