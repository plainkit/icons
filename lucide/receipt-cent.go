package lucide

import (
	html "github.com/plainkit/html"
)

// ReceiptCent creates a Receipt Cent Lucide icon.
func ReceiptCent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt-cent", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z"))),
		html.Child(html.SvgPath(html.AD("M12 6.5v11"))),
		html.Child(html.SvgPath(html.AD("M15 9.4a4 4 0 1 0 0 5.2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
