package validators

import "regexp"

func IsEmailValid(email string) bool {
	rex := `^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$`
	return regexp.MustCompile(rex).MatchString(email)
}
