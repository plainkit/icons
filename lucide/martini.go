package lucide

import (
	html "github.com/plainkit/html"
)

// Martini creates a Martini Lucide icon.
func Martini(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-martini", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 22h8"))),
		html.Child(html.SvgPath(html.AD("M12 11v11"))),
		html.Child(html.SvgPath(html.AD("m19 3-7 8-7-8Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
