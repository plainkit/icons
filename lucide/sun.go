package lucide

import (
	html "github.com/plainkit/html"
)

// Sun creates a Sun Lucide icon.
func Sun(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sun", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4"))),
		html.Child(html.SvgPath(html.AD("M12 2v2"))),
		html.Child(html.SvgPath(html.AD("M12 20v2"))),
		html.Child(html.SvgPath(html.AD("m4.93 4.93 1.41 1.41"))),
		html.Child(html.SvgPath(html.AD("m17.66 17.66 1.41 1.41"))),
		html.Child(html.SvgPath(html.AD("M2 12h2"))),
		html.Child(html.SvgPath(html.AD("M20 12h2"))),
		html.Child(html.SvgPath(html.AD("m6.34 17.66-1.41 1.41"))),
		html.Child(html.SvgPath(html.AD("m19.07 4.93-1.41 1.41"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
