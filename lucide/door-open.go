package lucide

import (
	html "github.com/plainkit/html"
)

// DoorOpen creates a Door Open Lucide icon.
func DoorOpen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-door-open", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 20H2")),
		html.SvgPath(html.AD("M11 4.562v16.157a1 1 0 0 0 1.242.97L19 20V5.562a2 2 0 0 0-1.515-1.94l-4-1A2 2 0 0 0 11 4.561z")),
		html.SvgPath(html.AD("M11 4H8a2 2 0 0 0-2 2v14")),
		html.SvgPath(html.AD("M14 12h.01")),
		html.SvgPath(html.AD("M22 20h-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
