package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeJapaneseYen creates a Badge Japanese Yen Lucide icon.
func BadgeJapaneseYen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-japanese-yen", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		html.Child(html.SvgPath(html.AD("m9 8 3 3v7"))),
		html.Child(html.SvgPath(html.AD("m12 11 3-3"))),
		html.Child(html.SvgPath(html.AD("M9 12h6"))),
		html.Child(html.SvgPath(html.AD("M9 16h6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
