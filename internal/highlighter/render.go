package highlighter

func Highlight(tokens []SemanticToken) string {
	var result string
	for _, t := range tokens {
		style := styleForToken(t)
		result += style.Render(t.Value) + " "
	}
	return result
}
