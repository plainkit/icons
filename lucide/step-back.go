package lucide

import x "github.com/plainkit/blox"

// StepBack creates a Step Back Lucide icon.
func StepBack(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-step-back", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.971 4.285A2 2 0 0 1 17 6v12a2 2 0 0 1-3.029 1.715l-9.997-5.998a2 2 0 0 1-.003-3.432z"))),
		x.Child(x.Path(x.D("M21 20V4"))),
	)
	return x.Svg(svgArgs...)
}
