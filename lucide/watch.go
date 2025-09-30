package lucide

import (
	html "github.com/plainkit/html"
)

// Watch creates a Watch Lucide icon.
func Watch(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-watch", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 10v2.2l1.6 1")),
		html.SvgPath(html.AD("m16.13 7.66-.81-4.05a2 2 0 0 0-2-1.61h-2.68a2 2 0 0 0-2 1.61l-.78 4.05")),
		html.SvgPath(html.AD("m7.88 16.36.8 4a2 2 0 0 0 2 1.61h2.72a2 2 0 0 0 2-1.61l.81-4.05")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
