package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeEuro creates a Badge Euro Lucide icon.
func BadgeEuro(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-euro", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z")),
		html.SvgPath(html.AD("M7 12h5")),
		html.SvgPath(html.AD("M15 9.4a4 4 0 1 0 0 5.2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
