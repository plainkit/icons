package lucide

import (
	html "github.com/plainkit/html"
)

// MouseOff creates a Mouse Off Lucide icon.
func MouseOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mouse-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 6v.343"))),
		html.Child(html.SvgPath(html.AD("M18.218 18.218A7 7 0 0 1 5 15V9a7 7 0 0 1 .782-3.218"))),
		html.Child(html.SvgPath(html.AD("M19 13.343V9A7 7 0 0 0 8.56 2.902"))),
		html.Child(html.SvgPath(html.AD("M22 22 2 2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
