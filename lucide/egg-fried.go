package lucide

import (
	html "github.com/plainkit/html"
)

// EggFried creates a Egg Fried Lucide icon.
func EggFried(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-egg-fried", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("11.5"), html.ACy("12.5"), html.AR("3.5")),
		html.SvgPath(html.AD("M3 8c0-3.5 2.5-6 6.5-6 5 0 4.83 3 7.5 5s5 2 5 6c0 4.5-2.5 6.5-7 6.5-2.5 0-2.5 2.5-6 2.5s-7-2-7-5.5c0-3 1.5-3 1.5-5C3.5 10 3 9 3 8Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
