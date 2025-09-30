package lucide

import (
	html "github.com/plainkit/html"
)

// LockKeyholeOpen creates a Lock Keyhole Open Lucide icon.
func LockKeyholeOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lock-keyhole-open", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("16"), html.AR("1")),
		html.SvgRect(html.AWidth("18"), html.AHeight("12"), html.AX("3"), html.AY("10"), html.ARx("2")),
		html.SvgPath(html.AD("M7 10V7a5 5 0 0 1 9.33-2.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
