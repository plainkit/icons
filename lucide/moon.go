package lucide

import (
	html "github.com/plainkit/html"
)

// Moon creates a Moon Lucide icon.
func Moon(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-moon", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M20.985 12.486a9 9 0 1 1-9.473-9.472c.405-.022.617.46.402.803a6 6 0 0 0 8.268 8.268c.344-.215.825-.004.803.401"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
