package lucide

import (
	html "github.com/plainkit/html"
)

// DoorClosedLocked creates a Door Closed Locked Lucide icon.
func DoorClosedLocked(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-door-closed-locked", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 12h.01"))),
		html.Child(html.SvgPath(html.AD("M18 9V6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14"))),
		html.Child(html.SvgPath(html.AD("M2 20h8"))),
		html.Child(html.SvgPath(html.AD("M20 17v-2a2 2 0 1 0-4 0v2"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("5"), html.AX("14"), html.AY("17"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
