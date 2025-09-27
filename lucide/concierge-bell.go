package lucide

import (
	html "github.com/plainkit/html"
)

// ConciergeBell creates a Concierge Bell Lucide icon.
func ConciergeBell(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-concierge-bell", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 20a1 1 0 0 1-1-1v-1a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1Z"))),
		html.Child(html.SvgPath(html.AD("M20 16a8 8 0 1 0-16 0"))),
		html.Child(html.SvgPath(html.AD("M12 4v4"))),
		html.Child(html.SvgPath(html.AD("M10 4h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
