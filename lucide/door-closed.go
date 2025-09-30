package lucide

import (
	html "github.com/plainkit/html"
)

// DoorClosed creates a Door Closed Lucide icon.
func DoorClosed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-door-closed", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 12h.01")),
		html.SvgPath(html.AD("M18 20V6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14")),
		html.SvgPath(html.AD("M2 20h20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
