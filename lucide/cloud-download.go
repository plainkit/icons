package lucide

import (
	html "github.com/plainkit/html"
)

// CloudDownload creates a Cloud Download Lucide icon.
func CloudDownload(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-download", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 13v8l-4-4")),
		html.SvgPath(html.AD("m12 21 4-4")),
		html.SvgPath(html.AD("M4.393 15.269A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.436 8.284")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
