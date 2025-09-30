package lucide

import (
	html "github.com/plainkit/html"
)

// DatabaseZap creates a Database Zap Lucide icon.
func DatabaseZap(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-database-zap", args)
	children := []html.SvgArg{
		html.SvgEllipse(html.ACx("12"), html.ACy("5"), html.ARx("9"), html.ARy("3")),
		html.SvgPath(html.AD("M3 5V19A9 3 0 0 0 15 21.84")),
		html.SvgPath(html.AD("M21 5V8")),
		html.SvgPath(html.AD("M21 12L18 17H22L19 22")),
		html.SvgPath(html.AD("M3 12A9 3 0 0 0 14.59 14.87")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
