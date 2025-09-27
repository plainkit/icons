package lucide

import (
	html "github.com/plainkit/html"
)

// SquarePlay creates a Square Play Lucide icon.
func SquarePlay(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-play", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M9 9.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997A1 1 0 0 1 9 14.996z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
