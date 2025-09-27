package lucide

import (
	html "github.com/plainkit/html"
)

// CalendarHeart creates a Calendar Heart Lucide icon.
func CalendarHeart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-calendar-heart", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12.127 22H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v5.125"))),
		html.Child(html.SvgPath(html.AD("M14.62 18.8A2.25 2.25 0 1 1 18 15.836a2.25 2.25 0 1 1 3.38 2.966l-2.626 2.856a.998.998 0 0 1-1.507 0z"))),
		html.Child(html.SvgPath(html.AD("M16 2v4"))),
		html.Child(html.SvgPath(html.AD("M3 10h18"))),
		html.Child(html.SvgPath(html.AD("M8 2v4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
