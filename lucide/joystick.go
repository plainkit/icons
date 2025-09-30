package lucide

import (
	html "github.com/plainkit/html"
)

// Joystick creates a Joystick Lucide icon.
func Joystick(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-joystick", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-2Z")),
		html.SvgPath(html.AD("M6 15v-2")),
		html.SvgPath(html.AD("M12 15V9")),
		html.SvgCircle(html.ACx("12"), html.ACy("6"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
