package lucide

import (
	html "github.com/plainkit/html"
)

// Clock2 creates a Clock 2 Lucide icon.
func Clock2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-2", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6v6l4-2")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
