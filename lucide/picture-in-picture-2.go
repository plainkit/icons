package lucide

import (
	html "github.com/plainkit/html"
)

// PictureInPicture2 creates a Picture In Picture 2 Lucide icon.
func PictureInPicture2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-picture-in-picture-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 9V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10c0 1.1.9 2 2 2h4"))),
		html.Child(html.SvgRect(html.AWidth("10"), html.AHeight("7"), html.AX("12"), html.AY("13"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
