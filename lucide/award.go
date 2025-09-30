package lucide

import (
	html "github.com/plainkit/html"
)

// Award creates a Award Lucide icon.
func Award(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-award", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m15.477 12.89 1.515 8.526a.5.5 0 0 1-.81.47l-3.58-2.687a1 1 0 0 0-1.197 0l-3.586 2.686a.5.5 0 0 1-.81-.469l1.514-8.526")),
		html.SvgCircle(html.ACx("12"), html.ACy("8"), html.AR("6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
