package lucide

import (
	html "github.com/plainkit/html"
)

// FileWarning creates a File Warning Lucide icon.
func FileWarning(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-warning", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		html.Child(html.SvgPath(html.AD("M12 9v4"))),
		html.Child(html.SvgPath(html.AD("M12 17h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
