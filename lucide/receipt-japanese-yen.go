package lucide

import (
	html "github.com/plainkit/html"
)

// ReceiptJapaneseYen creates a Receipt Japanese Yen Lucide icon.
func ReceiptJapaneseYen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt-japanese-yen", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z")),
		html.SvgPath(html.AD("m12 10 3-3")),
		html.SvgPath(html.AD("m9 7 3 3v7.5")),
		html.SvgPath(html.AD("M9 11h6")),
		html.SvgPath(html.AD("M9 15h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
