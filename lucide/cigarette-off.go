package lucide

import x "github.com/plainkit/blox"

// CigaretteOff creates a Cigarette Off Lucide icon.
func CigaretteOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cigarette-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 12H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h13"))),
		x.Child(x.Path(x.D("M18 8c0-2.5-2-2.5-2-5"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M21 12a1 1 0 0 1 1 1v2a1 1 0 0 1-.5.866"))),
		x.Child(x.Path(x.D("M22 8c0-2.5-2-2.5-2-5"))),
		x.Child(x.Path(x.D("M7 12v4"))),
	)
	return x.Svg(svgArgs...)
}
