package lucide

import (
	html "github.com/plainkit/html"
)

// FileClock creates a File Clock Lucide icon.
func FileClock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-clock", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M16 22h2a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3"))),
		html.Child(html.SvgPath(html.AD("M8 14v2.2l1.6 1"))),
		html.Child(html.SvgCircle(html.ACx("8"), html.ACy("16"), html.AR("6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
