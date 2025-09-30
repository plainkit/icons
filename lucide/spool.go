package lucide

import (
	html "github.com/plainkit/html"
)

// Spool creates a Spool Lucide icon.
func Spool(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-spool", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 13.44 4.442 17.082A2 2 0 0 0 4.982 21H19a2 2 0 0 0 .558-3.921l-1.115-.32A2 2 0 0 1 17 14.837V7.66")),
		html.SvgPath(html.AD("m7 10.56 12.558-3.642A2 2 0 0 0 19.018 3H5a2 2 0 0 0-.558 3.921l1.115.32A2 2 0 0 1 7 9.163v7.178")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
