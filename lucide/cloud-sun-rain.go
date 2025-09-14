package lucide

import x "github.com/bloxui/blox"

// CloudSunRain creates a Cloud Sun Rain Lucide icon.
func CloudSunRain(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-sun-rain", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("m4.93 4.93 1.41 1.41"))),
		x.Child(x.Path(x.D("M20 12h2"))),
		x.Child(x.Path(x.D("m19.07 4.93-1.41 1.41"))),
		x.Child(x.Path(x.D("M15.947 12.65a4 4 0 0 0-5.925-4.128"))),
		x.Child(x.Path(x.D("M3 20a5 5 0 1 1 8.9-4H13a3 3 0 0 1 2 5.24"))),
		x.Child(x.Path(x.D("M11 20v2"))),
		x.Child(x.Path(x.D("M7 19v2"))),
	)
	return x.Svg(svgArgs...)
}
