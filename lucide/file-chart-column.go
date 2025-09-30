package lucide

import (
	html "github.com/plainkit/html"
)

// FileChartColumn creates a File Chart Column Lucide icon.
func FileChartColumn(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-chart-column", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("M8 18v-1")),
		html.SvgPath(html.AD("M12 18v-6")),
		html.SvgPath(html.AD("M16 18v-3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
