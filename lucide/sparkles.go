package lucide

import (
	html "github.com/plainkit/html"
)

// Sparkles creates a Sparkles Lucide icon.
func Sparkles(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sparkles", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11.017 2.814a1 1 0 0 1 1.966 0l1.051 5.558a2 2 0 0 0 1.594 1.594l5.558 1.051a1 1 0 0 1 0 1.966l-5.558 1.051a2 2 0 0 0-1.594 1.594l-1.051 5.558a1 1 0 0 1-1.966 0l-1.051-5.558a2 2 0 0 0-1.594-1.594l-5.558-1.051a1 1 0 0 1 0-1.966l5.558-1.051a2 2 0 0 0 1.594-1.594z"))),
		html.Child(html.SvgPath(html.AD("M20 2v4"))),
		html.Child(html.SvgPath(html.AD("M22 4h-4"))),
		html.Child(html.SvgCircle(html.ACx("4"), html.ACy("20"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
