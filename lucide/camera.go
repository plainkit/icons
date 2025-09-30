package lucide

import (
	html "github.com/plainkit/html"
)

// Camera creates a Camera Lucide icon.
func Camera(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-camera", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13.997 4a2 2 0 0 1 1.76 1.05l.486.9A2 2 0 0 0 18.003 7H20a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1.997a2 2 0 0 0 1.759-1.048l.489-.904A2 2 0 0 1 10.004 4z")),
		html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
