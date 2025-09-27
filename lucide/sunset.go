package lucide

import (
	html "github.com/plainkit/html"
)

// Sunset creates a Sunset Lucide icon.
func Sunset(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sunset", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 10V2"))),
		html.Child(html.SvgPath(html.AD("m4.93 10.93 1.41 1.41"))),
		html.Child(html.SvgPath(html.AD("M2 18h2"))),
		html.Child(html.SvgPath(html.AD("M20 18h2"))),
		html.Child(html.SvgPath(html.AD("m19.07 10.93-1.41 1.41"))),
		html.Child(html.SvgPath(html.AD("M22 22H2"))),
		html.Child(html.SvgPath(html.AD("m16 6-4 4-4-4"))),
		html.Child(html.SvgPath(html.AD("M16 18a4 4 0 0 0-8 0"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
