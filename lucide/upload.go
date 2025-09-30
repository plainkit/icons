package lucide

import (
	html "github.com/plainkit/html"
)

// Upload creates a Upload Lucide icon.
func Upload(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-upload", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 3v12")),
		html.SvgPath(html.AD("m17 8-5-5-5 5")),
		html.SvgPath(html.AD("M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
