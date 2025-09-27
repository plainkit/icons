package lucide

import (
	html "github.com/plainkit/html"
)

// SmilePlus creates a Smile Plus Lucide icon.
func SmilePlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-smile-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 11v1a10 10 0 1 1-9-10"))),
		html.Child(html.SvgPath(html.AD("M8 14s1.5 2 4 2 4-2 4-2"))),
		html.Child(html.SvgLine(html.AX1("9"), html.AX2("9.01"), html.AY1("9"), html.AY2("9"))),
		html.Child(html.SvgLine(html.AX1("15"), html.AX2("15.01"), html.AY1("9"), html.AY2("9"))),
		html.Child(html.SvgPath(html.AD("M16 5h6"))),
		html.Child(html.SvgPath(html.AD("M19 2v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
