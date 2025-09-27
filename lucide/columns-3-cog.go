package lucide

import (
	html "github.com/plainkit/html"
)

// Columns3Cog creates a Columns 3 Cog Lucide icon.
func Columns3Cog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-columns-3-cog", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.5 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v5.5"))),
		html.Child(html.SvgPath(html.AD("m14.3 19.6 1-.4"))),
		html.Child(html.SvgPath(html.AD("M15 3v7.5"))),
		html.Child(html.SvgPath(html.AD("m15.2 16.9-.9-.3"))),
		html.Child(html.SvgPath(html.AD("m16.6 21.7.3-.9"))),
		html.Child(html.SvgPath(html.AD("m16.8 15.3-.4-1"))),
		html.Child(html.SvgPath(html.AD("m19.1 15.2.3-.9"))),
		html.Child(html.SvgPath(html.AD("m19.6 21.7-.4-1"))),
		html.Child(html.SvgPath(html.AD("m20.7 16.8 1-.4"))),
		html.Child(html.SvgPath(html.AD("m21.7 19.4-.9-.3"))),
		html.Child(html.SvgPath(html.AD("M9 3v18"))),
		html.Child(html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
