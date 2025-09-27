package lucide

import (
	html "github.com/plainkit/html"
)

// Database creates a Database Lucide icon.
func Database(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-database", args)
	children := []html.SvgArg{
		html.Child(html.SvgEllipse(html.ACx("12"), html.ACy("5"), html.ARx("9"), html.ARy("3"))),
		html.Child(html.SvgPath(html.AD("M3 5V19A9 3 0 0 0 21 19V5"))),
		html.Child(html.SvgPath(html.AD("M3 12A9 3 0 0 0 21 12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
