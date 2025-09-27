package lucide

import (
	html "github.com/plainkit/html"
)

// Clock10 creates a Clock 10 Lucide icon.
func Clock10(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clock-10", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 6v6l-4-2"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
