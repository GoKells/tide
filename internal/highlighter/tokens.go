package highlighter

import (
	"regexp"
	"strings"
)

type Token struct {
	Start int    // start position in the file/text
	End   int    // end position
	Type  string // "keyword", "string", "number", etc.
}


func ParseSemanticTokens(input string) []SemanticToken {
	pattern := `/\*⇒(\d+),([^,]+),(\[[^\]]*\])\*/([^\s]+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	var tokens []SemanticToken
	for _, m := range matches {
		modifiers := strings.Trim(m[3], "[]")
		mods := []string{}
		if modifiers != "" {
			mods = strings.Split(modifiers, " ")
		}

		tokens = append(tokens, SemanticToken{
			Type:      m[2],
			Modifiers: mods,
			Value:     m[4],
		})
	}

	return tokens
}
