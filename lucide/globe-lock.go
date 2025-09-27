package lucide

import (
	html "github.com/plainkit/html"
)

// GlobeLock creates a Globe Lock Lucide icon.
func GlobeLock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-globe-lock", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15.686 15A14.5 14.5 0 0 1 12 22a14.5 14.5 0 0 1 0-20 10 10 0 1 0 9.542 13"))),
		html.Child(html.SvgPath(html.AD("M2 12h8.5"))),
		html.Child(html.SvgPath(html.AD("M20 6V4a2 2 0 1 0-4 0v2"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("5"), html.AX("14"), html.AY("6"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
