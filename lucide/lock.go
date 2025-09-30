package lucide

import (
	html "github.com/plainkit/html"
)

// Lock creates a Lock Lucide icon.
func Lock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lock", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("11"), html.AX("3"), html.AY("11"), html.ARx("2"), html.ARy("2")),
		html.SvgPath(html.AD("M7 11V7a5 5 0 0 1 10 0v4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
