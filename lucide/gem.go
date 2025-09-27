package lucide

import (
	html "github.com/plainkit/html"
)

// Gem creates a Gem Lucide icon.
func Gem(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gem", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.5 3 8 9l4 13 4-13-2.5-6"))),
		html.Child(html.SvgPath(html.AD("M17 3a2 2 0 0 1 1.6.8l3 4a2 2 0 0 1 .013 2.382l-7.99 10.986a2 2 0 0 1-3.247 0l-7.99-10.986A2 2 0 0 1 2.4 7.8l2.998-3.997A2 2 0 0 1 7 3z"))),
		html.Child(html.SvgPath(html.AD("M2 9h20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
