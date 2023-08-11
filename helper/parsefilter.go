package helper

import "strings"

func ParseFilter(filter string) string {

	trim := strings.TrimSpace(filter)

	split := strings.Split(trim, " ")

	join := strings.Join(split, "* AND ")

	join += "* AND"

	return token(join)
}

func token(data string) string {
	phrase := data
	suffix := " AND"

	lastIndex := strings.LastIndex(phrase, suffix)

	if lastIndex != -1 {
		updatedPhrase := phrase[:lastIndex] + phrase[lastIndex+len(suffix):]
		return updatedPhrase
	}

	return phrase
}
