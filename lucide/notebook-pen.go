package lucide

import (
	html "github.com/plainkit/html"
)

// NotebookPen creates a Notebook Pen Lucide icon.
func NotebookPen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-notebook-pen", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13.4 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-7.4")),
		html.SvgPath(html.AD("M2 6h4")),
		html.SvgPath(html.AD("M2 10h4")),
		html.SvgPath(html.AD("M2 14h4")),
		html.SvgPath(html.AD("M2 18h4")),
		html.SvgPath(html.AD("M21.378 5.626a1 1 0 1 0-3.004-3.004l-5.01 5.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
