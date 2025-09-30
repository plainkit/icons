package lucide

import (
	html "github.com/plainkit/html"
)

// Clock8 creates a Clock 8 Lucide icon.
func Clock8(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-8", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6v6l-4 2")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
