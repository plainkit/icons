package lucide

import (
	html "github.com/plainkit/html"
)

// Clock9 creates a Clock 9 Lucide icon.
func Clock9(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-9", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 6v6H8")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
