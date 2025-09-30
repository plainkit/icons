package lucide

import (
	html "github.com/plainkit/html"
)

// School creates a School Lucide icon.
func School(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-school", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 21v-3a2 2 0 0 0-4 0v3")),
		html.SvgPath(html.AD("M18 5v16")),
		html.SvgPath(html.AD("m4 6 7.106-3.79a2 2 0 0 1 1.788 0L20 6")),
		html.SvgPath(html.AD("m6 11-3.52 2.147a1 1 0 0 0-.48.854V19a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5a1 1 0 0 0-.48-.853L18 11")),
		html.SvgPath(html.AD("M6 5v16")),
		html.SvgCircle(html.ACx("12"), html.ACy("9"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
