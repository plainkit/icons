package lucide

import (
	html "github.com/plainkit/html"
)

// WifiCog creates a Wifi Cog Lucide icon.
func WifiCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wifi-cog", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m14.305 19.53.923-.382"))),
		html.Child(html.SvgPath(html.AD("m15.228 16.852-.923-.383"))),
		html.Child(html.SvgPath(html.AD("m16.852 15.228-.383-.923"))),
		html.Child(html.SvgPath(html.AD("m16.852 20.772-.383.924"))),
		html.Child(html.SvgPath(html.AD("m19.148 15.228.383-.923"))),
		html.Child(html.SvgPath(html.AD("m19.53 21.696-.382-.924"))),
		html.Child(html.SvgPath(html.AD("M2 7.82a15 15 0 0 1 20 0"))),
		html.Child(html.SvgPath(html.AD("m20.772 16.852.924-.383"))),
		html.Child(html.SvgPath(html.AD("m20.772 19.148.924.383"))),
		html.Child(html.SvgPath(html.AD("M5 11.858a10 10 0 0 1 11.5-1.785"))),
		html.Child(html.SvgPath(html.AD("M8.5 15.429a5 5 0 0 1 2.413-1.31"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
