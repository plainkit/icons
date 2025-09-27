package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeRussianRuble creates a Badge Russian Ruble Lucide icon.
func BadgeRussianRuble(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-russian-ruble", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		html.Child(html.SvgPath(html.AD("M9 16h5"))),
		html.Child(html.SvgPath(html.AD("M9 12h5a2 2 0 1 0 0-4h-3v9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
