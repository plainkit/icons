package lucide

import (
	html "github.com/plainkit/html"
)

// EqualApproximately creates a Equal Approximately Lucide icon.
func EqualApproximately(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-equal-approximately", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 15a6.5 6.5 0 0 1 7 0 6.5 6.5 0 0 0 7 0"))),
		html.Child(html.SvgPath(html.AD("M5 9a6.5 6.5 0 0 1 7 0 6.5 6.5 0 0 0 7 0"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
