package lucide

import (
	html "github.com/plainkit/html"
)

// SaveOff creates a Save Off Lucide icon.
func SaveOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-save-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 13H8a1 1 0 0 0-1 1v7"))),
		html.Child(html.SvgPath(html.AD("M14 8h1"))),
		html.Child(html.SvgPath(html.AD("M17 21v-4"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M20.41 20.41A2 2 0 0 1 19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 .59-1.41"))),
		html.Child(html.SvgPath(html.AD("M29.5 11.5s5 5 4 5"))),
		html.Child(html.SvgPath(html.AD("M9 3h6.2a2 2 0 0 1 1.4.6l3.8 3.8a2 2 0 0 1 .6 1.4V15"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
