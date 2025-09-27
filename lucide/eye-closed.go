package lucide

import (
	html "github.com/plainkit/html"
)

// EyeClosed creates a Eye Closed Lucide icon.
func EyeClosed(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-eye-closed", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m15 18-.722-3.25"))),
		html.Child(html.SvgPath(html.AD("M2 8a10.645 10.645 0 0 0 20 0"))),
		html.Child(html.SvgPath(html.AD("m20 15-1.726-2.05"))),
		html.Child(html.SvgPath(html.AD("m4 15 1.726-2.05"))),
		html.Child(html.SvgPath(html.AD("m9 18 .722-3.25"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
