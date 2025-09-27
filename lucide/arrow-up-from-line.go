package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUpFromLine creates a Arrow Up From Line Lucide icon.
func ArrowUpFromLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-from-line", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m18 9-6-6-6 6"))),
		html.Child(html.SvgPath(html.AD("M12 3v14"))),
		html.Child(html.SvgPath(html.AD("M5 21h14"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
