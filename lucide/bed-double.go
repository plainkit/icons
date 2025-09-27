package lucide

import (
	html "github.com/plainkit/html"
)

// BedDouble creates a Bed Double Lucide icon.
func BedDouble(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bed-double", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 20v-8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8"))),
		html.Child(html.SvgPath(html.AD("M4 10V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4"))),
		html.Child(html.SvgPath(html.AD("M12 4v6"))),
		html.Child(html.SvgPath(html.AD("M2 18h20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
