package lucide

import (
	html "github.com/plainkit/html"
)

// FileStack creates a File Stack Lucide icon.
func FileStack(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-stack", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 21a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1"))),
		html.Child(html.SvgPath(html.AD("M16 16a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1"))),
		html.Child(html.SvgPath(html.AD("M21 6a2 2 0 0 0-.586-1.414l-2-2A2 2 0 0 0 17 2h-3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
