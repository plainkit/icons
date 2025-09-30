package lucide

import (
	html "github.com/plainkit/html"
)

// Cookie creates a Cookie Lucide icon.
func Cookie(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cookie", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 2a10 10 0 1 0 10 10 4 4 0 0 1-5-5 4 4 0 0 1-5-5")),
		html.SvgPath(html.AD("M8.5 8.5v.01")),
		html.SvgPath(html.AD("M16 15.5v.01")),
		html.SvgPath(html.AD("M12 12v.01")),
		html.SvgPath(html.AD("M11 17v.01")),
		html.SvgPath(html.AD("M7 14v.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
