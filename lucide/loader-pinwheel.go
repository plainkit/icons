package lucide

import (
	html "github.com/plainkit/html"
)

// LoaderPinwheel creates a Loader Pinwheel Lucide icon.
func LoaderPinwheel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-loader-pinwheel", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 12a1 1 0 0 1-10 0 1 1 0 0 0-10 0")),
		html.SvgPath(html.AD("M7 20.7a1 1 0 1 1 5-8.7 1 1 0 1 0 5-8.6")),
		html.SvgPath(html.AD("M7 3.3a1 1 0 1 1 5 8.6 1 1 0 1 0 5 8.6")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
