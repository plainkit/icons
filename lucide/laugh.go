package lucide

import (
	html "github.com/plainkit/html"
)

// Laugh creates a Laugh Lucide icon.
func Laugh(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-laugh", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M18 13a6 6 0 0 1-6 5 6 6 0 0 1-6-5h12Z"))),
		html.Child(html.SvgLine(html.AX1("9"), html.AX2("9.01"), html.AY1("9"), html.AY2("9"))),
		html.Child(html.SvgLine(html.AX1("15"), html.AX2("15.01"), html.AY1("9"), html.AY2("9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
