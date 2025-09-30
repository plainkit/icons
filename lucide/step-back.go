package lucide

import (
	html "github.com/plainkit/html"
)

// StepBack creates a Step Back Lucide icon.
func StepBack(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-step-back", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13.971 4.285A2 2 0 0 1 17 6v12a2 2 0 0 1-3.029 1.715l-9.997-5.998a2 2 0 0 1-.003-3.432z")),
		html.SvgPath(html.AD("M21 20V4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
