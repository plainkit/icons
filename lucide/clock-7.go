package lucide

import (
	html "github.com/plainkit/html"
)

// Clock7 creates a Clock 7 Lucide icon.
func Clock7(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-7", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 6v6l-2 4"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
