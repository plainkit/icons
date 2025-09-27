package lucide

import (
	html "github.com/plainkit/html"
)

// SquareSigma creates a Square Sigma Lucide icon.
func SquareSigma(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-sigma", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M16 8.9V7H8l4 5-4 5h8v-1.9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
