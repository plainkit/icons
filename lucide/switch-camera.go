package lucide

import (
	html "github.com/plainkit/html"
)

// SwitchCamera creates a Switch Camera Lucide icon.
func SwitchCamera(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-switch-camera", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 19H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h5")),
		html.SvgPath(html.AD("M13 5h7a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-5")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3")),
		html.SvgPath(html.AD("m18 22-3-3 3-3")),
		html.SvgPath(html.AD("m6 2 3 3-3 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
