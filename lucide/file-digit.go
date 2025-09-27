package lucide

import (
	html "github.com/plainkit/html"
)

// FileDigit creates a File Digit Lucide icon.
func FileDigit(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-digit", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("6"), html.AX("2"), html.AY("12"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M10 12h2v6"))),
		html.Child(html.SvgPath(html.AD("M10 18h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
