package utilities

import "unicode"

func CamelToSnakeCase(input string) string {
	var result []rune

	for i, char := range input {
		if i > 0 && unicode.IsUpper(char) {
			// If it's an uppercase letter and not the first character,
			// add an underscore before appending the lowercase version of the letter.
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(char))
	}

	return string(result)
}
