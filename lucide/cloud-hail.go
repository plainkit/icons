package lucide

import (
	html "github.com/plainkit/html"
)

// CloudHail creates a Cloud Hail Lucide icon.
func CloudHail(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-hail", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"))),
		html.Child(html.SvgPath(html.AD("M16 14v2"))),
		html.Child(html.SvgPath(html.AD("M8 14v2"))),
		html.Child(html.SvgPath(html.AD("M16 20h.01"))),
		html.Child(html.SvgPath(html.AD("M8 20h.01"))),
		html.Child(html.SvgPath(html.AD("M12 16v2"))),
		html.Child(html.SvgPath(html.AD("M12 22h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
