package lucide

import (
	html "github.com/plainkit/html"
)

// UserRoundCog creates a User Round Cog Lucide icon.
func UserRoundCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-round-cog", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m14.305 19.53.923-.382"))),
		html.Child(html.SvgPath(html.AD("m15.228 16.852-.923-.383"))),
		html.Child(html.SvgPath(html.AD("m16.852 15.228-.383-.923"))),
		html.Child(html.SvgPath(html.AD("m16.852 20.772-.383.924"))),
		html.Child(html.SvgPath(html.AD("m19.148 15.228.383-.923"))),
		html.Child(html.SvgPath(html.AD("m19.53 21.696-.382-.924"))),
		html.Child(html.SvgPath(html.AD("M2 21a8 8 0 0 1 10.434-7.62"))),
		html.Child(html.SvgPath(html.AD("m20.772 16.852.924-.383"))),
		html.Child(html.SvgPath(html.AD("m20.772 19.148.924.383"))),
		html.Child(html.SvgCircle(html.ACx("10"), html.ACy("8"), html.AR("5"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
