package lucide

import (
	html "github.com/plainkit/html"
)

// Disc creates a Disc Lucide icon.
func Disc(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-disc", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
