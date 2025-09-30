package lucide

import (
	html "github.com/plainkit/html"
)

// Sigma creates a Sigma Lucide icon.
func Sigma(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sigma", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 7V5a1 1 0 0 0-1-1H6.5a.5.5 0 0 0-.4.8l4.5 6a2 2 0 0 1 0 2.4l-4.5 6a.5.5 0 0 0 .4.8H17a1 1 0 0 0 1-1v-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
