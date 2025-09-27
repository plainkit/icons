package lucide

import (
	html "github.com/plainkit/html"
)

// MapMinus creates a Map Minus Lucide icon.
func MapMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-map-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m11 19-1.106-.552a2 2 0 0 0-1.788 0l-3.659 1.83A1 1 0 0 1 3 19.381V6.618a1 1 0 0 1 .553-.894l4.553-2.277a2 2 0 0 1 1.788 0l4.212 2.106a2 2 0 0 0 1.788 0l3.659-1.83A1 1 0 0 1 21 4.619V14"))),
		html.Child(html.SvgPath(html.AD("M15 5.764V14"))),
		html.Child(html.SvgPath(html.AD("M21 18h-6"))),
		html.Child(html.SvgPath(html.AD("M9 3.236v15"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
