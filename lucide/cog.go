package lucide

import (
	html "github.com/plainkit/html"
)

// Cog creates a Cog Lucide icon.
func Cog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cog", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 10.27 7 3.34")),
		html.SvgPath(html.AD("m11 13.73-4 6.93")),
		html.SvgPath(html.AD("M12 22v-2")),
		html.SvgPath(html.AD("M12 2v2")),
		html.SvgPath(html.AD("M14 12h8")),
		html.SvgPath(html.AD("m17 20.66-1-1.73")),
		html.SvgPath(html.AD("m17 3.34-1 1.73")),
		html.SvgPath(html.AD("M2 12h2")),
		html.SvgPath(html.AD("m20.66 17-1.73-1")),
		html.SvgPath(html.AD("m20.66 7-1.73 1")),
		html.SvgPath(html.AD("m3.34 17 1.73-1")),
		html.SvgPath(html.AD("m3.34 7 1.73 1")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
