package flotilla

import "regexp"

func Components(re *regexp.Regexp, path string) map[string]string {
	names := re.SubexpNames()
	match := re.FindStringSubmatch(path)
	if match == nil {
		return nil
	}
	m := map[string]string{}
	for i, s := range match {
		m[names[i]] = s
	}
	return m
}
