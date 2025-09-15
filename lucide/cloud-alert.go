package lucide

import x "github.com/plainkit/html"

// CloudAlert creates a Cloud Alert Lucide icon.
func CloudAlert(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-alert", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 12v4"))),
		x.Child(x.Path(x.D("M12 20h.01"))),
		x.Child(x.Path(x.D("M17 18h.5a1 1 0 0 0 0-9h-1.79A7 7 0 1 0 7 17.708"))),
	)
	return x.Svg(svgArgs...)
}
