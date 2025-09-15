package lucide

import x "github.com/plainkit/blox"

// LampCeiling creates a Lamp Ceiling Lucide icon.
func LampCeiling(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-lamp-ceiling", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v5"))),
		x.Child(x.Path(x.D("M14.829 15.998a3 3 0 1 1-5.658 0"))),
		x.Child(x.Path(x.D("M20.92 14.606A1 1 0 0 1 20 16H4a1 1 0 0 1-.92-1.394l3-7A1 1 0 0 1 7 7h10a1 1 0 0 1 .92.606z"))),
	)
	return x.Svg(svgArgs...)
}
