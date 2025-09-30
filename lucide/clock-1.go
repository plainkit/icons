package lucide

import (
	html "github.com/plainkit/html"
)

// Clock1 creates a Clock 1 Lucide icon.
func Clock1(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-1", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6v6l2-4")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
