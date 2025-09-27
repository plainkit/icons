package lucide

import (
	html "github.com/plainkit/html"
)

// FileKey2 creates a File Key 2 Lucide icon.
func FileKey2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-key-2", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v6"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgCircle(html.ACx("4"), html.ACy("16"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("m10 10-4.5 4.5"))),
		html.Child(html.SvgPath(html.AD("m9 11 1 1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
