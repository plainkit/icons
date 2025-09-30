package lucide

import (
	html "github.com/plainkit/html"
)

// Eclipse creates a Eclipse Lucide icon.
func Eclipse(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-eclipse", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M12 2a7 7 0 1 0 10 10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
