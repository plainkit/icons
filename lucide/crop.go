package lucide

import (
	html "github.com/plainkit/html"
)

// Crop creates a Crop Lucide icon.
func Crop(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-crop", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 2v14a2 2 0 0 0 2 2h14"))),
		html.Child(html.SvgPath(html.AD("M18 22V8a2 2 0 0 0-2-2H2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
