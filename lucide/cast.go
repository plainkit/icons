package lucide

import (
	html "github.com/plainkit/html"
)

// Cast creates a Cast Lucide icon.
func Cast(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cast", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 8V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-6"))),
		html.Child(html.SvgPath(html.AD("M2 12a9 9 0 0 1 8 8"))),
		html.Child(html.SvgPath(html.AD("M2 16a5 5 0 0 1 4 4"))),
		html.Child(html.SvgLine(html.AX1("2"), html.AX2("2.01"), html.AY1("20"), html.AY2("20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
