package lucide

import (
	html "github.com/plainkit/html"
)

// Parentheses creates a Parentheses Lucide icon.
func Parentheses(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-parentheses", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8 21s-4-3-4-9 4-9 4-9")),
		html.SvgPath(html.AD("M16 3s4 3 4 9-4 9-4 9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
