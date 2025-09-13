package lucide

import x "github.com/bloxui/blox"

// ShieldAlert creates a Shield Alert Lucide icon.
func ShieldAlert(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-shield-alert", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z"))),
		x.Child(x.Path(x.D("M12 8v4"))),
		x.Child(x.Path(x.D("M12 16h.01"))),
	)
	return x.Svg(svgArgs...)
}
