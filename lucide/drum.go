package lucide

import (
	html "github.com/plainkit/html"
)

// Drum creates a Drum Lucide icon.
func Drum(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-drum", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m2 2 8 8"))),
		html.Child(html.SvgPath(html.AD("m22 2-8 8"))),
		html.Child(html.SvgEllipse(html.ACx("12"), html.ACy("9"), html.ARx("10"), html.ARy("5"))),
		html.Child(html.SvgPath(html.AD("M7 13.4v7.9"))),
		html.Child(html.SvgPath(html.AD("M12 14v8"))),
		html.Child(html.SvgPath(html.AD("M17 13.4v7.9"))),
		html.Child(html.SvgPath(html.AD("M2 9v8a10 5 0 0 0 20 0V9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
