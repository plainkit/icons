package lucide

import (
	html "github.com/plainkit/html"
)

// FileChartPie creates a File Chart Pie Lucide icon.
func FileChartPie(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-chart-pie", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("M16 22h2a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3.5"))),
		html.Child(html.SvgPath(html.AD("M4.017 11.512a6 6 0 1 0 8.466 8.475"))),
		html.Child(html.SvgPath(html.AD("M9 16a1 1 0 0 1-1-1v-4c0-.552.45-1.008.995-.917a6 6 0 0 1 4.922 4.922c.091.544-.365.995-.917.995z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
