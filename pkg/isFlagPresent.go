package pkg

import "flag"

func IsFlagPresent (name string) bool {
	present := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			present = true
		}
	})
	return present
}