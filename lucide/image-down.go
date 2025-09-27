package lucide

import (
	html "github.com/plainkit/html"
)

// ImageDown creates a Image Down Lucide icon.
func ImageDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-image-down", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.3 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10l-3.1-3.1a2 2 0 0 0-2.814.014L6 21"))),
		html.Child(html.SvgPath(html.AD("m14 19 3 3v-5.5"))),
		html.Child(html.SvgPath(html.AD("m17 22 3-3"))),
		html.Child(html.SvgCircle(html.ACx("9"), html.ACy("9"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
