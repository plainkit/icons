package lucide

import (
	html "github.com/plainkit/html"
)

// ImagePlay creates a Image Play Lucide icon.
func ImagePlay(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-image-play", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 15.003a1 1 0 0 1 1.517-.859l4.997 2.997a1 1 0 0 1 0 1.718l-4.997 2.997a1 1 0 0 1-1.517-.86z"))),
		html.Child(html.SvgPath(html.AD("M21 12.17V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"))),
		html.Child(html.SvgPath(html.AD("m6 21 5-5"))),
		html.Child(html.SvgCircle(html.ACx("9"), html.ACy("9"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
