package lucide

import (
	html "github.com/plainkit/html"
)

// PictureInPicture creates a Picture In Picture Lucide icon.
func PictureInPicture(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-picture-in-picture", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 10h6V4")),
		html.SvgPath(html.AD("m2 4 6 6")),
		html.SvgPath(html.AD("M21 10V7a2 2 0 0 0-2-2h-7")),
		html.SvgPath(html.AD("M3 14v2a2 2 0 0 0 2 2h3")),
		html.SvgRect(html.AWidth("10"), html.AHeight("7"), html.AX("12"), html.AY("14"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
