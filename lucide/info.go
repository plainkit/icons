package lucide

import (
	html "github.com/plainkit/html"
)

// Info creates a Info Lucide icon.
func Info(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-info", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M12 16v-4"))),
		html.Child(html.SvgPath(html.AD("M12 8h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
