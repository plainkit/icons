package lucide

import (
	html "github.com/plainkit/html"
)

// Disc2 creates a Disc 2 Lucide icon.
func Disc2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-disc-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4"))),
		html.Child(html.SvgPath(html.AD("M12 12h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
