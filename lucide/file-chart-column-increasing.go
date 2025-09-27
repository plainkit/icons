package lucide

import (
	html "github.com/plainkit/html"
)

// FileChartColumnIncreasing creates a File Chart Column Increasing Lucide icon.
func FileChartColumnIncreasing(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-chart-column-increasing", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M8 18v-2"))),
		html.Child(html.SvgPath(html.AD("M12 18v-4"))),
		html.Child(html.SvgPath(html.AD("M16 18v-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
