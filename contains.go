package main

func contains(str string, s []string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
