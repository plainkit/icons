package lucide

import (
	html "github.com/plainkit/html"
)

// Thermometer creates a Thermometer Lucide icon.
func Thermometer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-thermometer", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
