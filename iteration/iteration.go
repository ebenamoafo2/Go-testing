package iteration

import "strings"



func Repeat(character string, count int) string {
	var repeated strings.Builder

	if count <= 0 {
		return ""
	}

	for i := 0; i < count; i++ {
		repeated.WriteString(character)
		
	}
	
	return repeated.String()// used to .String method to retrieve the final string from the builder
}
