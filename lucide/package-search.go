package lucide

import (
	html "github.com/plainkit/html"
)

// PackageSearch creates a Package Search Lucide icon.
func PackageSearch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-package-search", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 10V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14")),
		html.SvgPath(html.AD("m7.5 4.27 9 5.15")),
		html.SvgPolyline(html.APoints("3.29 7 12 12 20.71 7")),
		html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("22"), html.AY2("12")),
		html.SvgCircle(html.ACx("18.5"), html.ACy("15.5"), html.AR("2.5")),
		html.SvgPath(html.AD("M20.27 17.27 22 19")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
