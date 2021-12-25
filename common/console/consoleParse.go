package console

import "regexp"

// ParseFlags parses input in the form of -key=value and returns the input as a map.
// Overwrites map values from the given `defaults` map, if a key match is found.
func ParseFlags(args []string, defaults map[string]string) map[string]string {
	var regexFlag = regexp.MustCompile(`^-+(\w+(?:-\w+)*)=([\W\w]+(?:-[\W\w]+)*)$`)

	for _, arg := range args {
		matches := regexFlag.FindStringSubmatch(arg)

		if len(matches) >= 2 {
			defaults[matches[1]] = matches[2]
		}
	}

	return defaults
}
