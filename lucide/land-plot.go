package lucide

import (
	html "github.com/plainkit/html"
)

// LandPlot creates a Land Plot Lucide icon.
func LandPlot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-land-plot", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m12 8 6-3-6-3v10")),
		html.SvgPath(html.AD("m8 11.99-5.5 3.14a1 1 0 0 0 0 1.74l8.5 4.86a2 2 0 0 0 2 0l8.5-4.86a1 1 0 0 0 0-1.74L16 12")),
		html.SvgPath(html.AD("m6.49 12.85 11.02 6.3")),
		html.SvgPath(html.AD("M17.51 12.85 6.5 19.15")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
