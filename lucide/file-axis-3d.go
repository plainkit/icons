package lucide

import (
	html "github.com/plainkit/html"
)

// FileAxis3d creates a File Axis 3d Lucide icon.
func FileAxis3d(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-axis-3d", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("m8 18 4-4"))),
		html.Child(html.SvgPath(html.AD("M8 10v8h8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
