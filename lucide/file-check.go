package lucide

import (
	html "github.com/plainkit/html"
)

// FileCheck creates a File Check Lucide icon.
func FileCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("m9 15 2 2 4-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
