package lucide

import (
	html "github.com/plainkit/html"
)

// Brackets creates a Brackets Lucide icon.
func Brackets(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-brackets", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 3h3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-3"))),
		html.Child(html.SvgPath(html.AD("M8 21H5a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
