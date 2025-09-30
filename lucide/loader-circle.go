package lucide

import (
	html "github.com/plainkit/html"
)

// LoaderCircle creates a Loader Circle Lucide icon.
func LoaderCircle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-loader-circle", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 12a9 9 0 1 1-6.219-8.56")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
