package lucide

import (
	html "github.com/plainkit/html"
)

// UserLock creates a User Lock Lucide icon.
func UserLock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-lock", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("10"), html.ACy("7"), html.AR("4")),
		html.SvgPath(html.AD("M10.3 15H7a4 4 0 0 0-4 4v2")),
		html.SvgPath(html.AD("M15 15.5V14a2 2 0 0 1 4 0v1.5")),
		html.SvgRect(html.AWidth("8"), html.AHeight("5"), html.AX("13"), html.AY("16"), html.ARx(".899")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
