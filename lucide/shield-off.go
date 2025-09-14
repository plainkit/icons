package lucide

import x "github.com/bloxui/blox"

// ShieldOff creates a Shield Off Lucide icon.
func ShieldOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-shield-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M5 5a1 1 0 0 0-1 1v7c0 5 3.5 7.5 7.67 8.94a1 1 0 0 0 .67.01c2.35-.82 4.48-1.97 5.9-3.71"))),
		x.Child(x.Path(x.D("M9.309 3.652A12.252 12.252 0 0 0 11.24 2.28a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1v7a9.784 9.784 0 0 1-.08 1.264"))),
	)
	return x.Svg(svgArgs...)
}
