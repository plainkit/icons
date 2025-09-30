package lucide

import (
	html "github.com/plainkit/html"
)

// Rotate3d creates a Rotate 3d Lucide icon.
func Rotate3d(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-rotate-3d", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M16.466 7.5C15.643 4.237 13.952 2 12 2 9.239 2 7 6.477 7 12s2.239 10 5 10c.342 0 .677-.069 1-.2")),
		html.SvgPath(html.AD("m15.194 13.707 3.814 1.86-1.86 3.814")),
		html.SvgPath(html.AD("M19 15.57c-1.804.885-4.274 1.43-7 1.43-5.523 0-10-2.239-10-5s4.477-5 10-5c4.838 0 8.873 1.718 9.8 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
