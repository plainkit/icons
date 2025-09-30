package lucide

import (
	html "github.com/plainkit/html"
)

// Venus creates a Venus Lucide icon.
func Venus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-venus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 15v7")),
		html.SvgPath(html.AD("M9 19h6")),
		html.SvgCircle(html.ACx("12"), html.ACy("9"), html.AR("6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
