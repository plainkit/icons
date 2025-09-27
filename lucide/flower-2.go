package lucide

import (
	html "github.com/plainkit/html"
)

// Flower2 creates a Flower 2 Lucide icon.
func Flower2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flower-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 5a3 3 0 1 1 3 3m-3-3a3 3 0 1 0-3 3m3-3v1M9 8a3 3 0 1 0 3 3M9 8h1m5 0a3 3 0 1 1-3 3m3-3h-1m-2 3v-1"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("8"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("M12 10v12"))),
		html.Child(html.SvgPath(html.AD("M12 22c4.2 0 7-1.667 7-5-4.2 0-7 1.667-7 5Z"))),
		html.Child(html.SvgPath(html.AD("M12 22c-4.2 0-7-1.667-7-5 4.2 0 7 1.667 7 5Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
