package lucide

import (
	html "github.com/plainkit/html"
)

// CloudCog creates a Cloud Cog Lucide icon.
func CloudCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-cog", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m10.852 19.772-.383.924")),
		html.SvgPath(html.AD("m13.148 14.228.383-.923")),
		html.SvgPath(html.AD("M13.148 19.772a3 3 0 1 0-2.296-5.544l-.383-.923")),
		html.SvgPath(html.AD("m13.53 20.696-.382-.924a3 3 0 1 1-2.296-5.544")),
		html.SvgPath(html.AD("m14.772 15.852.923-.383")),
		html.SvgPath(html.AD("m14.772 18.148.923.383")),
		html.SvgPath(html.AD("M4.2 15.1a7 7 0 1 1 9.93-9.858A7 7 0 0 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.2")),
		html.SvgPath(html.AD("m9.228 15.852-.923-.383")),
		html.SvgPath(html.AD("m9.228 18.148-.923.383")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
