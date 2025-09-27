package lucide

import (
	html "github.com/plainkit/html"
)

// Mailbox creates a Mailbox Lucide icon.
func Mailbox(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mailbox", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M22 17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9.5C2 7 4 5 6.5 5H18c2.2 0 4 1.8 4 4v8Z"))),
		html.Child(html.SvgPolyline(html.APoints("15,9 18,9 18,11"))),
		html.Child(html.SvgPath(html.AD("M6.5 5C9 5 11 7 11 9.5V17a2 2 0 0 1-2 2"))),
		html.Child(html.SvgLine(html.AX1("6"), html.AX2("7"), html.AY1("10"), html.AY2("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
