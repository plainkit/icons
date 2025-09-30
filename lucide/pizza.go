package lucide

import (
	html "github.com/plainkit/html"
)

// Pizza creates a Pizza Lucide icon.
func Pizza(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-pizza", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m12 14-1 1")),
		html.SvgPath(html.AD("m13.75 18.25-1.25 1.42")),
		html.SvgPath(html.AD("M17.775 5.654a15.68 15.68 0 0 0-12.121 12.12")),
		html.SvgPath(html.AD("M18.8 9.3a1 1 0 0 0 2.1 7.7")),
		html.SvgPath(html.AD("M21.964 20.732a1 1 0 0 1-1.232 1.232l-18-5a1 1 0 0 1-.695-1.232A19.68 19.68 0 0 1 15.732 2.037a1 1 0 0 1 1.232.695z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
