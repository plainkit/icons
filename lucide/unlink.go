package lucide

import (
	html "github.com/plainkit/html"
)

// Unlink creates a Unlink Lucide icon.
func Unlink(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-unlink", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m18.84 12.25 1.72-1.71h-.02a5.004 5.004 0 0 0-.12-7.07 5.006 5.006 0 0 0-6.95 0l-1.72 1.71"))),
		html.Child(html.SvgPath(html.AD("m5.17 11.75-1.71 1.71a5.004 5.004 0 0 0 .12 7.07 5.006 5.006 0 0 0 6.95 0l1.71-1.71"))),
		html.Child(html.SvgLine(html.AX1("8"), html.AX2("8"), html.AY1("2"), html.AY2("5"))),
		html.Child(html.SvgLine(html.AX1("2"), html.AX2("5"), html.AY1("8"), html.AY2("8"))),
		html.Child(html.SvgLine(html.AX1("16"), html.AX2("16"), html.AY1("19"), html.AY2("22"))),
		html.Child(html.SvgLine(html.AX1("19"), html.AX2("22"), html.AY1("16"), html.AY2("16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
