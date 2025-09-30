package lucide

import (
	html "github.com/plainkit/html"
)

// Radio creates a Radio Lucide icon.
func Radio(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-radio", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16.247 7.761a6 6 0 0 1 0 8.478")),
		html.SvgPath(html.AD("M19.075 4.933a10 10 0 0 1 0 14.134")),
		html.SvgPath(html.AD("M4.925 19.067a10 10 0 0 1 0-14.134")),
		html.SvgPath(html.AD("M7.753 16.239a6 6 0 0 1 0-8.478")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
