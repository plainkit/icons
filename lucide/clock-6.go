package lucide

import (
	html "github.com/plainkit/html"
)

// Clock6 creates a Clock 6 Lucide icon.
func Clock6(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-6", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6v10")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
