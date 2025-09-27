package lucide

import (
	html "github.com/plainkit/html"
)

// FileCog creates a File Cog Lucide icon.
func FileCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-file-cog", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 2v4a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgPath(html.AD("m2.305 15.53.923-.382"))),
		html.Child(html.SvgPath(html.AD("m3.228 12.852-.924-.383"))),
		html.Child(html.SvgPath(html.AD("M4.677 21.5a2 2 0 0 0 1.313.5H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v2.5"))),
		html.Child(html.SvgPath(html.AD("m4.852 11.228-.383-.923"))),
		html.Child(html.SvgPath(html.AD("m4.852 16.772-.383.924"))),
		html.Child(html.SvgPath(html.AD("m7.148 11.228.383-.923"))),
		html.Child(html.SvgPath(html.AD("m7.53 17.696-.382-.924"))),
		html.Child(html.SvgPath(html.AD("m8.772 12.852.923-.383"))),
		html.Child(html.SvgPath(html.AD("m8.772 15.148.923.383"))),
		html.Child(html.SvgCircle(html.ACx("6"), html.ACy("14"), html.AR("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
