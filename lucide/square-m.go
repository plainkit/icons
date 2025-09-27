package lucide

import (
	html "github.com/plainkit/html"
)

// SquareM creates a Square M Lucide icon.
func SquareM(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-m", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8 16V8.5a.5.5 0 0 1 .9-.3l2.7 3.599a.5.5 0 0 0 .8 0l2.7-3.6a.5.5 0 0 1 .9.3V16"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
