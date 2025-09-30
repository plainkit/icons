package lucide

import (
	html "github.com/plainkit/html"
)

// CloudDrizzle creates a Cloud Drizzle Lucide icon.
func CloudDrizzle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-drizzle", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242")),
		html.SvgPath(html.AD("M8 19v1")),
		html.SvgPath(html.AD("M8 14v1")),
		html.SvgPath(html.AD("M16 19v1")),
		html.SvgPath(html.AD("M16 14v1")),
		html.SvgPath(html.AD("M12 21v1")),
		html.SvgPath(html.AD("M12 16v1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
