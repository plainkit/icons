package lucide

import (
	html "github.com/plainkit/html"
)

// CloudLightning creates a Cloud Lightning Lucide icon.
func CloudLightning(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-lightning", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 16.326A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 .5 8.973")),
		html.SvgPath(html.AD("m13 12-3 5h4l-3 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
