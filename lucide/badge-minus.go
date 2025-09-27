package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeMinus creates a Badge Minus Lucide icon.
func BadgeMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-minus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		html.Child(html.SvgLine(html.AX1("8"), html.AX2("16"), html.AY1("12"), html.AY2("12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
