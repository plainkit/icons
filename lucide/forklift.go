package lucide

import (
	html "github.com/plainkit/html"
)

// Forklift creates a Forklift Lucide icon.
func Forklift(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-forklift", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 12H5a2 2 0 0 0-2 2v5")),
		html.SvgCircle(html.ACx("13"), html.ACy("19"), html.AR("2")),
		html.SvgCircle(html.ACx("5"), html.ACy("19"), html.AR("2")),
		html.SvgPath(html.AD("M8 19h3m5-17v17h6M6 12V7c0-1.1.9-2 2-2h3l5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
