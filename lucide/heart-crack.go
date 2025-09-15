package lucide

import x "github.com/plainkit/html"

// HeartCrack creates a Heart Crack Lucide icon.
func HeartCrack(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-heart-crack", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.409 5.824c-.702.792-1.15 1.496-1.415 2.166l2.153 2.156a.5.5 0 0 1 0 .707l-2.293 2.293a.5.5 0 0 0 0 .707L12 15"))),
		x.Child(x.Path(x.D("M13.508 20.313a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5a5.5 5.5 0 0 1 9.591-3.677.6.6 0 0 0 .818.001A5.5 5.5 0 0 1 22 9.5c0 2.29-1.5 4-3 5.5z"))),
	)
	return x.Svg(svgArgs...)
}
