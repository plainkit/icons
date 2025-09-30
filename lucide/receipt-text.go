package lucide

import (
	html "github.com/plainkit/html"
)

// ReceiptText creates a Receipt Text Lucide icon.
func ReceiptText(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt-text", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z")),
		html.SvgPath(html.AD("M14 8H8")),
		html.SvgPath(html.AD("M16 12H8")),
		html.SvgPath(html.AD("M13 16H8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
