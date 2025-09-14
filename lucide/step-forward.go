package lucide

import x "github.com/bloxui/blox"

// StepForward creates a Step Forward Lucide icon.
func StepForward(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-step-forward", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.029 4.285A2 2 0 0 0 7 6v12a2 2 0 0 0 3.029 1.715l9.997-5.998a2 2 0 0 0 .003-3.432z"))),
		x.Child(x.Path(x.D("M3 4v16"))),
	)
	return x.Svg(svgArgs...)
}
