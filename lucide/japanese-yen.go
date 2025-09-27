package lucide

import (
	html "github.com/plainkit/html"
)

// JapaneseYen creates a Japanese Yen Lucide icon.
func JapaneseYen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-japanese-yen", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 9.5V21m0-11.5L6 3m6 6.5L18 3"))),
		html.Child(html.SvgPath(html.AD("M6 15h12"))),
		html.Child(html.SvgPath(html.AD("M6 11h12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
