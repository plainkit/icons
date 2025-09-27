package lucide

import (
	html "github.com/plainkit/html"
)

// CloudSun creates a Cloud Sun Lucide icon.
func CloudSun(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-sun", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2v2"))),
		html.Child(html.SvgPath(html.AD("m4.93 4.93 1.41 1.41"))),
		html.Child(html.SvgPath(html.AD("M20 12h2"))),
		html.Child(html.SvgPath(html.AD("m19.07 4.93-1.41 1.41"))),
		html.Child(html.SvgPath(html.AD("M15.947 12.65a4 4 0 0 0-5.925-4.128"))),
		html.Child(html.SvgPath(html.AD("M13 22H7a5 5 0 1 1 4.9-6H13a3 3 0 0 1 0 6Z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
