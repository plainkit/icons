package lucide

import x "github.com/bloxui/blox"

// CloudSun creates a Cloud Sun Lucide icon.
func CloudSun(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-sun", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("m4.93 4.93 1.41 1.41"))),
		x.Child(x.Path(x.D("M20 12h2"))),
		x.Child(x.Path(x.D("m19.07 4.93-1.41 1.41"))),
		x.Child(x.Path(x.D("M15.947 12.65a4 4 0 0 0-5.925-4.128"))),
		x.Child(x.Path(x.D("M13 22H7a5 5 0 1 1 4.9-6H13a3 3 0 0 1 0 6Z"))),
	)
	return x.Svg(svgArgs...)
}
