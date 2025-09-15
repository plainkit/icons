package lucide

import x "github.com/plainkit/blox"

// CloudOff creates a Cloud Off Lucide icon.
func CloudOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M5.782 5.782A7 7 0 0 0 9 19h8.5a4.5 4.5 0 0 0 1.307-.193"))),
		x.Child(x.Path(x.D("M21.532 16.5A4.5 4.5 0 0 0 17.5 10h-1.79A7.008 7.008 0 0 0 10 5.07"))),
	)
	return x.Svg(svgArgs...)
}
