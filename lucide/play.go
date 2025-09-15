package lucide

import x "github.com/plainkit/blox"

// Play creates a Play Lucide icon.
func Play(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-play", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5 5a2 2 0 0 1 3.008-1.728l11.997 6.998a2 2 0 0 1 .003 3.458l-12 7A2 2 0 0 1 5 19z"))),
	)
	return x.Svg(svgArgs...)
}
