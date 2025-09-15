package lucide

import x "github.com/plainkit/blox"

// SaveOff creates a Save Off Lucide icon.
func SaveOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-save-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 13H8a1 1 0 0 0-1 1v7"))),
		x.Child(x.Path(x.D("M14 8h1"))),
		x.Child(x.Path(x.D("M17 21v-4"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M20.41 20.41A2 2 0 0 1 19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 .59-1.41"))),
		x.Child(x.Path(x.D("M29.5 11.5s5 5 4 5"))),
		x.Child(x.Path(x.D("M9 3h6.2a2 2 0 0 1 1.4.6l3.8 3.8a2 2 0 0 1 .6 1.4V15"))),
	)
	return x.Svg(svgArgs...)
}
