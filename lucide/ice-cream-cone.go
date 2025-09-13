package lucide

import x "github.com/bloxui/blox"

// IceCreamCone creates a Ice Cream Cone Lucide icon.
func IceCreamCone(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-ice-cream-cone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m7 11 4.08 10.35a1 1 0 0 0 1.84 0L17 11"))),
		x.Child(x.Path(x.D("M17 7A5 5 0 0 0 7 7"))),
		x.Child(x.Path(x.D("M17 7a2 2 0 0 1 0 4H7a2 2 0 0 1 0-4"))),
	)
	return x.Svg(svgArgs...)
}
