package lucide

import (
	html "github.com/plainkit/html"
)

// SwatchBook creates a Swatch Book Lucide icon.
func SwatchBook(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-swatch-book", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 17a4 4 0 0 1-8 0V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2Z")),
		html.SvgPath(html.AD("M16.7 13H19a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H7")),
		html.SvgPath(html.AD("M 7 17h.01")),
		html.SvgPath(html.AD("m11 8 2.3-2.3a2.4 2.4 0 0 1 3.404.004L18.6 7.6a2.4 2.4 0 0 1 .026 3.434L9.9 19.8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
