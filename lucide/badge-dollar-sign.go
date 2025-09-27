package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeDollarSign creates a Badge Dollar Sign Lucide icon.
func BadgeDollarSign(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-dollar-sign", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		html.Child(html.SvgPath(html.AD("M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8"))),
		html.Child(html.SvgPath(html.AD("M12 18V6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
