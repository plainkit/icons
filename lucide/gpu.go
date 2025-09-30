package lucide

import (
	html "github.com/plainkit/html"
)

// Gpu creates a Gpu Lucide icon.
func Gpu(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gpu", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 21V3")),
		html.SvgPath(html.AD("M2 5h18a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2.26")),
		html.SvgPath(html.AD("M7 17v3a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-3")),
		html.SvgCircle(html.ACx("16"), html.ACy("11"), html.AR("2")),
		html.SvgCircle(html.ACx("8"), html.ACy("11"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
