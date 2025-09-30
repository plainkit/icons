package lucide

import (
	html "github.com/plainkit/html"
)

// MoonStar creates a Moon Star Lucide icon.
func MoonStar(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-moon-star", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 5h4")),
		html.SvgPath(html.AD("M20 3v4")),
		html.SvgPath(html.AD("M20.985 12.486a9 9 0 1 1-9.473-9.472c.405-.022.617.46.402.803a6 6 0 0 0 8.268 8.268c.344-.215.825-.004.803.401")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
