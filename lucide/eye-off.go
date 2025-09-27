package lucide

import (
	html "github.com/plainkit/html"
)

// EyeOff creates a Eye Off Lucide icon.
func EyeOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-eye-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.733 5.076a10.744 10.744 0 0 1 11.205 6.575 1 1 0 0 1 0 .696 10.747 10.747 0 0 1-1.444 2.49"))),
		html.Child(html.SvgPath(html.AD("M14.084 14.158a3 3 0 0 1-4.242-4.242"))),
		html.Child(html.SvgPath(html.AD("M17.479 17.499a10.75 10.75 0 0 1-15.417-5.151 1 1 0 0 1 0-.696 10.75 10.75 0 0 1 4.446-5.143"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
