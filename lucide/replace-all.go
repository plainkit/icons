package lucide

import (
	html "github.com/plainkit/html"
)

// ReplaceAll creates a Replace All Lucide icon.
func ReplaceAll(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-replace-all", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 14a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2"))),
		html.Child(html.SvgPath(html.AD("M14 4a2 2 0 0 1 2-2"))),
		html.Child(html.SvgPath(html.AD("M16 10a2 2 0 0 1-2-2"))),
		html.Child(html.SvgPath(html.AD("M20 14a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2"))),
		html.Child(html.SvgPath(html.AD("M20 2a2 2 0 0 1 2 2"))),
		html.Child(html.SvgPath(html.AD("M22 8a2 2 0 0 1-2 2"))),
		html.Child(html.SvgPath(html.AD("m3 7 3 3 3-3"))),
		html.Child(html.SvgPath(html.AD("M6 10V5a 3 3 0 0 1 3-3h1"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("2"), html.AY("14"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
