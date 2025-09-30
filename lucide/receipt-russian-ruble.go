package lucide

import (
	html "github.com/plainkit/html"
)

// ReceiptRussianRuble creates a Receipt Russian Ruble Lucide icon.
func ReceiptRussianRuble(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt-russian-ruble", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z")),
		html.SvgPath(html.AD("M8 15h5")),
		html.SvgPath(html.AD("M8 11h5a2 2 0 1 0 0-4h-3v10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
