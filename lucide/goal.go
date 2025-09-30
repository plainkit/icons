package lucide

import (
	html "github.com/plainkit/html"
)

// Goal creates a Goal Lucide icon.
func Goal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-goal", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 13V2l8 4-8 4")),
		html.SvgPath(html.AD("M20.561 10.222a9 9 0 1 1-12.55-5.29")),
		html.SvgPath(html.AD("M8.002 9.997a5 5 0 1 0 8.9 2.02")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
