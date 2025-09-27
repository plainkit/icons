package lucide

import (
	html "github.com/plainkit/html"
)

// ReceiptTurkishLira creates a Receipt Turkish Lira Lucide icon.
func ReceiptTurkishLira(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt-turkish-lira", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 6.5v11a5.5 5.5 0 0 0 5.5-5.5"))),
		html.Child(html.SvgPath(html.AD("m14 8-6 3"))),
		html.Child(html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
