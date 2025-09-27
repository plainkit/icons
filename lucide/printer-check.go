package lucide

import (
	html "github.com/plainkit/html"
)

// PrinterCheck creates a Printer Check Lucide icon.
func PrinterCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-printer-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13.5 22H7a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v.5"))),
		html.Child(html.SvgPath(html.AD("m16 19 2 2 4-4"))),
		html.Child(html.SvgPath(html.AD("M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2"))),
		html.Child(html.SvgPath(html.AD("M6 9V3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
