package lucide

import x "github.com/bloxui/blox"

// BellRing creates a Bell Ring Lucide icon.
func BellRing(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-bell-ring", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.268 21a2 2 0 0 0 3.464 0"))),
		x.Child(x.Path(x.D("M22 8c0-2.3-.8-4.3-2-6"))),
		x.Child(x.Path(x.D("M3.262 15.326A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673C19.41 13.956 18 12.499 18 8A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.738 7.326"))),
		x.Child(x.Path(x.D("M4 2C2.8 3.7 2 5.7 2 8"))),
	)
	return x.Svg(svgArgs...)
}
