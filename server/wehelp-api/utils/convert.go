package utils

import "strconv"

// ConvertID : covnert ID string to ID int32.
func ConvertID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}
