package lucide

import (
	html "github.com/plainkit/html"
)

// HouseHeart creates a House Heart Lucide icon.
func HouseHeart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-house-heart", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M8.62 13.8A2.25 2.25 0 1 1 12 10.836a2.25 2.25 0 1 1 3.38 2.966l-2.626 2.856a.998.998 0 0 1-1.507 0z"))),
		html.Child(html.SvgPath(html.AD("M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
