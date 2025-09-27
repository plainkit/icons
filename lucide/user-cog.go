package lucide

import (
	html "github.com/plainkit/html"
)

// UserCog creates a User Cog Lucide icon.
func UserCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-user-cog", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 15H6a4 4 0 0 0-4 4v2"))),
		html.Child(html.SvgPath(html.AD("m14.305 16.53.923-.382"))),
		html.Child(html.SvgPath(html.AD("m15.228 13.852-.923-.383"))),
		html.Child(html.SvgPath(html.AD("m16.852 12.228-.383-.923"))),
		html.Child(html.SvgPath(html.AD("m16.852 17.772-.383.924"))),
		html.Child(html.SvgPath(html.AD("m19.148 12.228.383-.923"))),
		html.Child(html.SvgPath(html.AD("m19.53 18.696-.382-.924"))),
		html.Child(html.SvgPath(html.AD("m20.772 13.852.924-.383"))),
		html.Child(html.SvgPath(html.AD("m20.772 16.148.924.383"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("15"), html.AR("3"))),
		html.Child(html.SvgCircle(html.ACx("9"), html.ACy("7"), html.AR("4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
