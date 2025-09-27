package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeInfo creates a Badge Info Lucide icon.
func BadgeInfo(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-info", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("16"), html.AY2("12"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12.01"), html.AY1("8"), html.AY2("8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
