package utils

import "regexp"

var regexpEmail *regexp.Regexp

func init() {
	regexpEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
}

// ValidEmail : check is email vaild or not.
// TODO : check this https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835
// TODO : check this "gopkg.in/go-playground/validator.v9"
func ValidEmail(input string) bool {
	return regexpEmail.MatchString(input)
}
