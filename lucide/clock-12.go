package lucide

import (
	html "github.com/plainkit/html"
)

// Clock12 creates a Clock 12 Lucide icon.
func Clock12(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-12", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6v6")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
