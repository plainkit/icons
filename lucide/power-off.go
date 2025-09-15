package lucide

import x "github.com/plainkit/html"

// PowerOff creates a Power Off Lucide icon.
func PowerOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-power-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18.36 6.64A9 9 0 0 1 20.77 15"))),
		x.Child(x.Path(x.D("M6.16 6.16a9 9 0 1 0 12.68 12.68"))),
		x.Child(x.Path(x.D("M12 2v4"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
