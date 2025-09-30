package lucide

import (
	html "github.com/plainkit/html"
)

// Barcode creates a Barcode Lucide icon.
func Barcode(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-barcode", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 5v14")),
		html.SvgPath(html.AD("M8 5v14")),
		html.SvgPath(html.AD("M12 5v14")),
		html.SvgPath(html.AD("M17 5v14")),
		html.SvgPath(html.AD("M21 5v14")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
