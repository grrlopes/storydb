package helper

import "strings"

func ParseFilter(filter string) string {
	var addQuotes []string

	filter = strings.ReplaceAll(filter, "\"", "'")

	trim := strings.TrimSpace(filter)

	split := strings.Split(trim, " ")

	for _, v := range split {
		addQuotes = append(addQuotes, `"`+v+`"`)
	}

	join := strings.Join(addQuotes, "* AND ")

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
