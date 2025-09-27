package lucide

import (
	html "github.com/plainkit/html"
)

// ServerCog creates a Server Cog Lucide icon.
func ServerCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-server-cog", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m10.852 14.772-.383.923"))),
		html.Child(html.SvgPath(html.AD("M13.148 14.772a3 3 0 1 0-2.296-5.544l-.383-.923"))),
		html.Child(html.SvgPath(html.AD("m13.148 9.228.383-.923"))),
		html.Child(html.SvgPath(html.AD("m13.53 15.696-.382-.924a3 3 0 1 1-2.296-5.544"))),
		html.Child(html.SvgPath(html.AD("m14.772 10.852.923-.383"))),
		html.Child(html.SvgPath(html.AD("m14.772 13.148.923.383"))),
		html.Child(html.SvgPath(html.AD("M4.5 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-.5"))),
		html.Child(html.SvgPath(html.AD("M4.5 14H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-.5"))),
		html.Child(html.SvgPath(html.AD("M6 18h.01"))),
		html.Child(html.SvgPath(html.AD("M6 6h.01"))),
		html.Child(html.SvgPath(html.AD("m9.228 10.852-.923-.383"))),
		html.Child(html.SvgPath(html.AD("m9.228 13.148-.923.383"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
