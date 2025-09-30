package lucide

import (
	html "github.com/plainkit/html"
)

// Toilet creates a Toilet Lucide icon.
func Toilet(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-toilet", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M7 12h13a1 1 0 0 1 1 1 5 5 0 0 1-5 5h-.598a.5.5 0 0 0-.424.765l1.544 2.47a.5.5 0 0 1-.424.765H5.402a.5.5 0 0 1-.424-.765L7 18")),
		html.SvgPath(html.AD("M8 18a5 5 0 0 1-5-5V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
