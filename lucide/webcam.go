package lucide

import (
	html "github.com/plainkit/html"
)

// Webcam creates a Webcam Lucide icon.
func Webcam(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-webcam", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("8")),
		html.SvgCircle(html.ACx("12"), html.ACy("10"), html.AR("3")),
		html.SvgPath(html.AD("M7 22h10")),
		html.SvgPath(html.AD("M12 22v-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
