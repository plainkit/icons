package lucide

import (
	html "github.com/plainkit/html"
)

// ReceiptPoundSterling creates a Receipt Pound Sterling Lucide icon.
func ReceiptPoundSterling(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt-pound-sterling", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z")),
		html.SvgPath(html.AD("M8 13h5")),
		html.SvgPath(html.AD("M10 17V9.5a2.5 2.5 0 0 1 5 0")),
		html.SvgPath(html.AD("M8 17h7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
