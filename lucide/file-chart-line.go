package lucide

import (
	html "github.com/plainkit/html"
)

// FileChartLine creates a File Chart Line Lucide icon.
func FileChartLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-chart-line", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z")),
		html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4")),
		html.SvgPath(html.AD("m16 13-3.5 3.5-2-2L8 17")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
