package lucide

import (
	html "github.com/plainkit/html"
)

// FileUser creates a File User Lucide icon.
func FileUser(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-user", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M15 18a3 3 0 1 0-6 0"))),
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7z"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
