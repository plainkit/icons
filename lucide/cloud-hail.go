package lucide

import x "github.com/bloxui/blox"

// CloudHail creates a Cloud Hail Lucide icon.
func CloudHail(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-hail", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"))),
		x.Child(x.Path(x.D("M16 14v2"))),
		x.Child(x.Path(x.D("M8 14v2"))),
		x.Child(x.Path(x.D("M16 20h.01"))),
		x.Child(x.Path(x.D("M8 20h.01"))),
		x.Child(x.Path(x.D("M12 16v2"))),
		x.Child(x.Path(x.D("M12 22h.01"))),
	)
	return x.Svg(svgArgs...)
}
