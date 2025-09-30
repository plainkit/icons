package lucide

import (
	html "github.com/plainkit/html"
)

// Play creates a Play Lucide icon.
func Play(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-play", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 5a2 2 0 0 1 3.008-1.728l11.997 6.998a2 2 0 0 1 .003 3.458l-12 7A2 2 0 0 1 5 19z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
