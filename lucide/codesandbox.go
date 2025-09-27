package lucide

import (
	html "github.com/plainkit/html"
)

// Codesandbox creates a Codesandbox Lucide icon.
func Codesandbox(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-codesandbox", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"))),
		html.Child(html.SvgPolyline(html.APoints("7.5 4.21 12 6.81 16.5 4.21"))),
		html.Child(html.SvgPolyline(html.APoints("7.5 19.79 7.5 14.6 3 12"))),
		html.Child(html.SvgPolyline(html.APoints("21 12 16.5 14.6 16.5 19.79"))),
		html.Child(html.SvgPolyline(html.APoints("3.27 6.96 12 12.01 20.73 6.96"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("22.08"), html.AY2("12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
