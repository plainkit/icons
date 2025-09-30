package lucide

import (
	html "github.com/plainkit/html"
)

// CircleSlash2 creates a Circle Slash 2 Lucide icon.
func CircleSlash2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-slash-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 2 2 22")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
