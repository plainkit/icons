package lucide

import (
	html "github.com/plainkit/html"
)

// LampCeiling creates a Lamp Ceiling Lucide icon.
func LampCeiling(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lamp-ceiling", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 2v5"))),
		html.Child(html.SvgPath(html.AD("M14.829 15.998a3 3 0 1 1-5.658 0"))),
		html.Child(html.SvgPath(html.AD("M20.92 14.606A1 1 0 0 1 20 16H4a1 1 0 0 1-.92-1.394l3-7A1 1 0 0 1 7 7h10a1 1 0 0 1 .92.606z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
