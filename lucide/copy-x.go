package lucide

import (
	html "github.com/plainkit/html"
)

// CopyX creates a Copy X Lucide icon.
func CopyX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-copy-x", args)
	children := []html.SvgArg{
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("18"), html.AY1("12"), html.AY2("18"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("18"), html.AY1("18"), html.AY2("12"))),
		html.Child(html.SvgRect(html.AWidth("14"), html.AHeight("14"), html.AX("8"), html.AY("8"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgPath(html.AD("M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
