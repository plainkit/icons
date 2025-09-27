package lucide

import (
	html "github.com/plainkit/html"
)

// LogOut creates a Log Out Lucide icon.
func LogOut(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-log-out", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m16 17 5-5-5-5"))),
		html.Child(html.SvgPath(html.AD("M21 12H9"))),
		html.Child(html.SvgPath(html.AD("M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
