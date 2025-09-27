package lucide

import (
	html "github.com/plainkit/html"
)

// Loader creates a Loader Lucide icon.
func Loader(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-loader", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2v4"))),
		html.Child(html.SvgPath(html.AD("m16.2 7.8 2.9-2.9"))),
		html.Child(html.SvgPath(html.AD("M18 12h4"))),
		html.Child(html.SvgPath(html.AD("m16.2 16.2 2.9 2.9"))),
		html.Child(html.SvgPath(html.AD("M12 18v4"))),
		html.Child(html.SvgPath(html.AD("m4.9 19.1 2.9-2.9"))),
		html.Child(html.SvgPath(html.AD("M2 12h4"))),
		html.Child(html.SvgPath(html.AD("m4.9 4.9 2.9 2.9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
