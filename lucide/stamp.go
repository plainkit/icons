package lucide

import (
	html "github.com/plainkit/html"
)

// Stamp creates a Stamp Lucide icon.
func Stamp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-stamp", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 13V8.5C14 7 15 7 15 5a3 3 0 0 0-6 0c0 2 1 2 1 3.5V13")),
		html.SvgPath(html.AD("M20 15.5a2.5 2.5 0 0 0-2.5-2.5h-11A2.5 2.5 0 0 0 4 15.5V17a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1z")),
		html.SvgPath(html.AD("M5 22h14")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
