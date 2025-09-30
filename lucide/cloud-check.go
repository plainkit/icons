package lucide

import (
	html "github.com/plainkit/html"
)

// CloudCheck creates a Cloud Check Lucide icon.
func CloudCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m17 15-5.5 5.5L9 18")),
		html.SvgPath(html.AD("M5 17.743A7 7 0 1 1 15.71 10h1.79a4.5 4.5 0 0 1 1.5 8.742")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
