package lucide

import (
	html "github.com/plainkit/html"
)

// CakeSlice creates a Cake Slice Lucide icon.
func CakeSlice(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cake-slice", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 13H3")),
		html.SvgPath(html.AD("M16 17H3")),
		html.SvgPath(html.AD("m7.2 7.9-3.388 2.5A2 2 0 0 0 3 12.01V20a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-8.654c0-2-2.44-6.026-6.44-8.026a1 1 0 0 0-1.082.057L10.4 5.6")),
		html.SvgCircle(html.ACx("9"), html.ACy("7"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
