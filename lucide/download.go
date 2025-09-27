package lucide

import (
	html "github.com/plainkit/html"
)

// Download creates a Download Lucide icon.
func Download(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-download", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 15V3"))),
		html.Child(html.SvgPath(html.AD("M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"))),
		html.Child(html.SvgPath(html.AD("m7 10 5 5 5-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
