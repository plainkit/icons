package lucide

import (
	html "github.com/plainkit/html"
)

// BrainCog creates a Brain Cog Lucide icon.
func BrainCog(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-brain-cog", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m10.852 14.772-.383.923")),
		html.SvgPath(html.AD("m10.852 9.228-.383-.923")),
		html.SvgPath(html.AD("m13.148 14.772.382.924")),
		html.SvgPath(html.AD("m13.531 8.305-.383.923")),
		html.SvgPath(html.AD("m14.772 10.852.923-.383")),
		html.SvgPath(html.AD("m14.772 13.148.923.383")),
		html.SvgPath(html.AD("M17.598 6.5A3 3 0 1 0 12 5a3 3 0 0 0-5.63-1.446 3 3 0 0 0-.368 1.571 4 4 0 0 0-2.525 5.771")),
		html.SvgPath(html.AD("M17.998 5.125a4 4 0 0 1 2.525 5.771")),
		html.SvgPath(html.AD("M19.505 10.294a4 4 0 0 1-1.5 7.706")),
		html.SvgPath(html.AD("M4.032 17.483A4 4 0 0 0 11.464 20c.18-.311.892-.311 1.072 0a4 4 0 0 0 7.432-2.516")),
		html.SvgPath(html.AD("M4.5 10.291A4 4 0 0 0 6 18")),
		html.SvgPath(html.AD("M6.002 5.125a3 3 0 0 0 .4 1.375")),
		html.SvgPath(html.AD("m9.228 10.852-.923-.383")),
		html.SvgPath(html.AD("m9.228 13.148-.923.383")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
