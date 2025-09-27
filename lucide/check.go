package lucide

import (
	html "github.com/plainkit/html"
)

// Check creates a Check Lucide icon.
func Check(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20 6 9 17l-5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
