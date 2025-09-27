package lucide

import (
	html "github.com/plainkit/html"
)

// Key creates a Key Lucide icon.
func Key(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-key", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m15.5 7.5 2.3 2.3a1 1 0 0 0 1.4 0l2.1-2.1a1 1 0 0 0 0-1.4L19 4"))),
		html.Child(html.SvgPath(html.AD("m21 2-9.6 9.6"))),
		html.Child(html.SvgCircle(html.ACx("7.5"), html.ACy("15.5"), html.AR("5.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
