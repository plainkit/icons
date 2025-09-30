package lucide

import (
	html "github.com/plainkit/html"
)

// Cloudy creates a Cloudy Lucide icon.
func Cloudy(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloudy", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17.5 21H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z")),
		html.SvgPath(html.AD("M22 10a3 3 0 0 0-3-3h-2.207a5.502 5.502 0 0 0-10.702.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
