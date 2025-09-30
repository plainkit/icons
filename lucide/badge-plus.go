package lucide

import (
	html "github.com/plainkit/html"
)

// BadgePlus creates a Badge Plus Lucide icon.
func BadgePlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-plus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z")),
		html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("8"), html.AY2("16")),
		html.SvgLine(html.AX1("8"), html.AX2("16"), html.AY1("12"), html.AY2("12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
