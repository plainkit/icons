package lucide

import (
	html "github.com/plainkit/html"
)

// ReceiptIndianRupee creates a Receipt Indian Rupee Lucide icon.
func ReceiptIndianRupee(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-receipt-indian-rupee", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 2v20l2-1 2 1 2-1 2 1 2-1 2 1 2-1 2 1V2l-2 1-2-1-2 1-2-1-2 1-2-1-2 1Z")),
		html.SvgPath(html.AD("M8 7h8")),
		html.SvgPath(html.AD("M12 17.5 8 15h1a4 4 0 0 0 0-8")),
		html.SvgPath(html.AD("M8 11h8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
