package lucide

import (
	html "github.com/plainkit/html"
)

// Logs creates a Logs Lucide icon.
func Logs(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-logs", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3 5h1")),
		html.SvgPath(html.AD("M3 12h1")),
		html.SvgPath(html.AD("M3 19h1")),
		html.SvgPath(html.AD("M8 5h1")),
		html.SvgPath(html.AD("M8 12h1")),
		html.SvgPath(html.AD("M8 19h1")),
		html.SvgPath(html.AD("M13 5h8")),
		html.SvgPath(html.AD("M13 12h8")),
		html.SvgPath(html.AD("M13 19h8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
