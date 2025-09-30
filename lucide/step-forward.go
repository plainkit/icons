package lucide

import (
	html "github.com/plainkit/html"
)

// StepForward creates a Step Forward Lucide icon.
func StepForward(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-step-forward", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.029 4.285A2 2 0 0 0 7 6v12a2 2 0 0 0 3.029 1.715l9.997-5.998a2 2 0 0 0 .003-3.432z")),
		html.SvgPath(html.AD("M3 4v16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
