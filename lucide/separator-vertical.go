package lucide

import (
	html "github.com/plainkit/html"
)

// SeparatorVertical creates a Separator Vertical Lucide icon.
func SeparatorVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-separator-vertical", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 3v18"))),
		html.Child(html.SvgPath(html.AD("m16 16 4-4-4-4"))),
		html.Child(html.SvgPath(html.AD("m8 8-4 4 4 4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
