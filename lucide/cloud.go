package lucide

import (
	html "github.com/plainkit/html"
)

// Cloud creates a Cloud Lucide icon.
func Cloud(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17.5 19H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
