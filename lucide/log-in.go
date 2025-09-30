package lucide

import (
	html "github.com/plainkit/html"
)

// LogIn creates a Log In Lucide icon.
func LogIn(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-log-in", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m10 17 5-5-5-5")),
		html.SvgPath(html.AD("M15 12H3")),
		html.SvgPath(html.AD("M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
