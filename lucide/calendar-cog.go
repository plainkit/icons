package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarCog creates a Calendar Cog Lucide icon.
func CalendarCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-cog", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m15.228 16.852-.923-.383"))),
		html.Child(html.SvgPath(html.AD("m15.228 19.148-.923.383"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgPath(html.AD("m16.47 14.305.382.923"))),
		html.Child(html.SvgPath(html.AD("m16.852 20.772-.383.924"))),
		html.Child(html.SvgPath(html.AD("m19.148 15.228.383-.923"))),
		html.Child(html.SvgPath(html.AD("m19.53 21.696-.382-.924"))),
		html.Child(html.SvgPath(html.AD("m20.772 16.852.924-.383"))),
		html.Child(html.SvgPath(html.AD("m20.772 19.148.924.383"))),
		html.Child(html.SvgPath(html.AD("M21 10.592V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
