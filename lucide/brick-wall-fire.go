package lucide

import (
	html "github.com/plainkit/html"
)

// BrickWallFire creates a Brick Wall Fire Lucide icon.
func BrickWallFire(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-brick-wall-fire", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 3v2.107"))),
		html.Child(html.SvgPath(html.AD("M17 9c1 3 2.5 3.5 3.5 4.5A5 5 0 0 1 22 17a5 5 0 0 1-10 0c0-.3 0-.6.1-.9a2 2 0 1 0 3.3-2C13 11.5 16 9 17 9"))),
		html.Child(html.SvgPath(html.AD("M21 8.274V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h3.938"))),
		html.Child(html.SvgPath(html.AD("M3 15h5.253"))),
		html.Child(html.SvgPath(html.AD("M3 9h8.228"))),
		html.Child(html.SvgPath(html.AD("M8 15v6"))),
		html.Child(html.SvgPath(html.AD("M8 3v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
