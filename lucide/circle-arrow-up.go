package lucide

import (
	html "github.com/plainkit/html"
)

// CircleArrowUp creates a Circle Arrow Up Lucide icon.
func CircleArrowUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-arrow-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("m16 12-4-4-4 4"))),
		html.Child(html.SvgPath(html.AD("M12 16V8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
