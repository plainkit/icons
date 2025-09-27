package lucide

import (
	html "github.com/plainkit/html"
)

// Bean creates a Bean Lucide icon.
func Bean(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bean", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.165 6.598C9.954 7.478 9.64 8.36 9 9c-.64.64-1.521.954-2.402 1.165A6 6 0 0 0 8 22c7.732 0 14-6.268 14-14a6 6 0 0 0-11.835-1.402Z"))),
		html.Child(html.SvgPath(html.AD("M5.341 10.62a4 4 0 1 0 5.279-5.28"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
