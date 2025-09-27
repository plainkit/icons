package lucide

import (
	html "github.com/plainkit/html"
)

// FlagOff creates a Flag Off Lucide icon.
func FlagOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flag-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 16c-3 0-5-2-8-2a6 6 0 0 0-4 1.528"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M4 22V4"))),
		html.Child(html.SvgPath(html.AD("M7.656 2H8c3 0 5 2 7.333 2q2 0 3.067-.8A1 1 0 0 1 20 4v10.347"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
