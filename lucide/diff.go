package lucide

import (
	html "github.com/plainkit/html"
)

// Diff creates a Diff Lucide icon.
func Diff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-diff", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 3v14"))),
		html.Child(html.SvgPath(html.AD("M5 10h14"))),
		html.Child(html.SvgPath(html.AD("M5 21h14"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
