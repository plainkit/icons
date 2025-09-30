package lucide

import (
	html "github.com/plainkit/html"
)

// CloudUpload creates a Cloud Upload Lucide icon.
func CloudUpload(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-upload", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 13v8")),
		html.SvgPath(html.AD("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242")),
		html.SvgPath(html.AD("m8 17 4-4 4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
