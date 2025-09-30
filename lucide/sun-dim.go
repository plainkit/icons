package lucide

import (
	html "github.com/plainkit/html"
)

// SunDim creates a Sun Dim Lucide icon.
func SunDim(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sun-dim", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4")),
		html.SvgPath(html.AD("M12 4h.01")),
		html.SvgPath(html.AD("M20 12h.01")),
		html.SvgPath(html.AD("M12 20h.01")),
		html.SvgPath(html.AD("M4 12h.01")),
		html.SvgPath(html.AD("M17.657 6.343h.01")),
		html.SvgPath(html.AD("M17.657 17.657h.01")),
		html.SvgPath(html.AD("M6.343 17.657h.01")),
		html.SvgPath(html.AD("M6.343 6.343h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
