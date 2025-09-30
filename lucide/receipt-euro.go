package lucide

import (
	html "github.com/plainkit/html"
)

// ReceiptEuro creates a Receipt Euro Lucide icon.
func ReceiptEuro(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt-euro", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z")),
		html.SvgPath(html.AD("M8 12h5")),
		html.SvgPath(html.AD("M16 9.5a4 4 0 1 0 0 5.2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
