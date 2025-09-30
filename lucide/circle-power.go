package lucide

import (
	html "github.com/plainkit/html"
)

// CirclePower creates a Circle Power Lucide icon.
func CirclePower(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-power", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 7v4")),
		html.SvgPath(html.AD("M7.998 9.003a5 5 0 1 0 8-.005")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
