package lucide

import (
	html "github.com/plainkit/html"
)

// Coins creates a Coins Lucide icon.
func Coins(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-coins", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("8"), html.ACy("8"), html.AR("6"))),
		html.Child(html.SvgPath(html.AD("M18.09 10.37A6 6 0 1 1 10.34 18"))),
		html.Child(html.SvgPath(html.AD("M7 6h1v4"))),
		html.Child(html.SvgPath(html.AD("m16.71 13.88.7.71-2.82 2.82"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
