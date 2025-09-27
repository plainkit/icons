package lucide

import (
	html "github.com/plainkit/html"
)

// PackagePlus creates a Package Plus Lucide icon.
func PackagePlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-package-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 16h6"))),
		html.Child(html.SvgPath(html.AD("M19 13v6"))),
		html.Child(html.SvgPath(html.AD("M21 10V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14"))),
		html.Child(html.SvgPath(html.AD("m7.5 4.27 9 5.15"))),
		html.Child(html.SvgPolyline(html.APoints("3.29 7 12 12 20.71 7"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("22"), html.AY2("12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
