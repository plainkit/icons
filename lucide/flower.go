package lucide

import (
	html "github.com/plainkit/html"
)

// Flower creates a Flower Lucide icon.
func Flower(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flower", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3"))),
		html.Child(html.SvgPath(html.AD("M12 16.5A4.5 4.5 0 1 1 7.5 12 4.5 4.5 0 1 1 12 7.5a4.5 4.5 0 1 1 4.5 4.5 4.5 4.5 0 1 1-4.5 4.5"))),
		html.Child(html.SvgPath(html.AD("M12 7.5V9"))),
		html.Child(html.SvgPath(html.AD("M7.5 12H9"))),
		html.Child(html.SvgPath(html.AD("M16.5 12H15"))),
		html.Child(html.SvgPath(html.AD("M12 16.5V15"))),
		html.Child(html.SvgPath(html.AD("m8 8 1.88 1.88"))),
		html.Child(html.SvgPath(html.AD("M14.12 9.88 16 8"))),
		html.Child(html.SvgPath(html.AD("m8 16 1.88-1.88"))),
		html.Child(html.SvgPath(html.AD("M14.12 14.12 16 16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
