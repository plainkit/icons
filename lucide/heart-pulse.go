package lucide

import x "github.com/bloxui/blox"

// HeartPulse creates a Heart Pulse Lucide icon.
func HeartPulse(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-heart-pulse", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 9.5a5.5 5.5 0 0 1 9.591-3.676.56.56 0 0 0 .818 0A5.49 5.49 0 0 1 22 9.5c0 2.29-1.5 4-3 5.5l-5.492 5.313a2 2 0 0 1-3 .019L5 15c-1.5-1.5-3-3.2-3-5.5"))),
		x.Child(x.Path(x.D("M3.22 13H9.5l.5-1 2 4.5 2-7 1.5 3.5h5.27"))),
	)
	return x.Svg(svgArgs...)
}
