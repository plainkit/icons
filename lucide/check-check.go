package lucide

import (
	html "github.com/plainkit/html"
)

// CheckCheck creates a Check Check Lucide icon.
func CheckCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-check-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 6 7 17l-5-5")),
		html.SvgPath(html.AD("m22 10-7.5 7.5L13 16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
