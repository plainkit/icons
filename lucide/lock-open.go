package lucide

import (
	html "github.com/plainkit/html"
)

// LockOpen creates a Lock Open Lucide icon.
func LockOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lock-open", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("11"), html.AX("3"), html.AY("11"), html.ARx("2"), html.ARy("2"))),
		html.Child(html.SvgPath(html.AD("M7 11V7a5 5 0 0 1 9.9-1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
