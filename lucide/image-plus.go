package lucide

import (
	html "github.com/plainkit/html"
)

// ImagePlus creates a Image Plus Lucide icon.
func ImagePlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-image-plus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16 5h6")),
		html.SvgPath(html.AD("M19 2v6")),
		html.SvgPath(html.AD("M21 11.5V19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7.5")),
		html.SvgPath(html.AD("m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21")),
		html.SvgCircle(html.ACx("9"), html.ACy("9"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
