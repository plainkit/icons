package lucide

import x "github.com/plainkit/blox"

// SkipForward creates a Skip Forward Lucide icon.
func SkipForward(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-skip-forward", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 4v16"))),
		x.Child(x.Path(x.D("M6.029 4.285A2 2 0 0 0 3 6v12a2 2 0 0 0 3.029 1.715l9.997-5.998a2 2 0 0 0 .003-3.432z"))),
	)
	return x.Svg(svgArgs...)
}
