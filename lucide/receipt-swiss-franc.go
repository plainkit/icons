package lucide

import (
	html "github.com/plainkit/html"
)

// ReceiptSwissFranc creates a Receipt Swiss Franc Lucide icon.
func ReceiptSwissFranc(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt-swiss-franc", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		html.Child(html.SvgPath(html.AD("M10 17V7h5"))),
		html.Child(html.SvgPath(html.AD("M10 11h4"))),
		html.Child(html.SvgPath(html.AD("M8 15h5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
