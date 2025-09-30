package lucide

import (
	html "github.com/plainkit/html"
)

// FolderCog creates a Folder Cog Lucide icon.
func FolderCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-folder-cog", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.3 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.98a2 2 0 0 1 1.69.9l.66 1.2A2 2 0 0 0 12 6h8a2 2 0 0 1 2 2v3.3")),
		html.SvgPath(html.AD("m14.305 19.53.923-.382")),
		html.SvgPath(html.AD("m15.228 16.852-.923-.383")),
		html.SvgPath(html.AD("m16.852 15.228-.383-.923")),
		html.SvgPath(html.AD("m16.852 20.772-.383.924")),
		html.SvgPath(html.AD("m19.148 15.228.383-.923")),
		html.SvgPath(html.AD("m19.53 21.696-.382-.924")),
		html.SvgPath(html.AD("m20.772 16.852.924-.383")),
		html.SvgPath(html.AD("m20.772 19.148.924.383")),
		html.SvgCircle(html.ACx("18"), html.ACy("18"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
