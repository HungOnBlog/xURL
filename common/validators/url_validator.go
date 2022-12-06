package validators

import "regexp"

func IsUrlValid(url string) bool {
	reg := `^((https?|ftp|smtp):\/\/)?(www.)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
	return regexp.MustCompile(reg).MatchString(url)
}
