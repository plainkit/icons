package lucide

import (
	html "github.com/plainkit/html"
)

// OctagonAlert creates a Octagon Alert Lucide icon.
func OctagonAlert(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-octagon-alert", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 16h.01")),
		html.SvgPath(html.AD("M12 8v4")),
		html.SvgPath(html.AD("M15.312 2a2 2 0 0 1 1.414.586l4.688 4.688A2 2 0 0 1 22 8.688v6.624a2 2 0 0 1-.586 1.414l-4.688 4.688a2 2 0 0 1-1.414.586H8.688a2 2 0 0 1-1.414-.586l-4.688-4.688A2 2 0 0 1 2 15.312V8.688a2 2 0 0 1 .586-1.414l4.688-4.688A2 2 0 0 1 8.688 2z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
