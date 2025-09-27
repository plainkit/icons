package lucide

import (
	html "github.com/plainkit/html"
)

// Figma creates a Figma Lucide icon.
func Figma(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-figma", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 5.5A3.5 3.5 0 0 1 8.5 2H12v7H8.5A3.5 3.5 0 0 1 5 5.5z"))),
		html.Child(html.SvgPath(html.AD("M12 2h3.5a3.5 3.5 0 1 1 0 7H12V2z"))),
		html.Child(html.SvgPath(html.AD("M12 12.5a3.5 3.5 0 1 1 7 0 3.5 3.5 0 1 1-7 0z"))),
		html.Child(html.SvgPath(html.AD("M5 19.5A3.5 3.5 0 0 1 8.5 16H12v3.5a3.5 3.5 0 1 1-7 0z"))),
		html.Child(html.SvgPath(html.AD("M5 12.5A3.5 3.5 0 0 1 8.5 9H12v7H8.5A3.5 3.5 0 0 1 5 12.5z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
