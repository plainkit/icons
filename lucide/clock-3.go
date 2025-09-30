package lucide

import (
	html "github.com/plainkit/html"
)

// Clock3 creates a Clock 3 Lucide icon.
func Clock3(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-3", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6v6h4")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
