package lucide

import (
	html "github.com/plainkit/html"
)

// CloudFog creates a Cloud Fog Lucide icon.
func CloudFog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-fog", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"))),
		html.Child(html.SvgPath(html.AD("M16 17H7"))),
		html.Child(html.SvgPath(html.AD("M17 21H9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
