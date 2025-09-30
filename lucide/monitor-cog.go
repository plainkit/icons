package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorCog creates a Monitor Cog Lucide icon.
func MonitorCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-cog", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 17v4")),
		html.SvgPath(html.AD("m14.305 7.53.923-.382")),
		html.SvgPath(html.AD("m15.228 4.852-.923-.383")),
		html.SvgPath(html.AD("m16.852 3.228-.383-.924")),
		html.SvgPath(html.AD("m16.852 8.772-.383.923")),
		html.SvgPath(html.AD("m19.148 3.228.383-.924")),
		html.SvgPath(html.AD("m19.53 9.696-.382-.924")),
		html.SvgPath(html.AD("m20.772 4.852.924-.383")),
		html.SvgPath(html.AD("m20.772 7.148.924.383")),
		html.SvgPath(html.AD("M22 13v2a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7")),
		html.SvgPath(html.AD("M8 21h8")),
		html.SvgCircle(html.ACx("18"), html.ACy("6"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
