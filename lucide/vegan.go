package lucide

import (
	html "github.com/plainkit/html"
)

// Vegan creates a Vegan Lucide icon.
func Vegan(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-vegan", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 8q6 0 6-6-6 0-6 6")),
		html.SvgPath(html.AD("M17.41 3.59a10 10 0 1 0 3 3")),
		html.SvgPath(html.AD("M2 2a26.6 26.6 0 0 1 10 20c.9-6.82 1.5-9.5 4-14")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
