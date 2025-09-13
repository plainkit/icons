package lucide

import x "github.com/bloxui/blox"

// Siren creates a Siren Lucide icon.
func Siren(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-siren", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 18v-6a5 5 0 1 1 10 0v6"))),
		x.Child(x.Path(x.D("M5 21a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2z"))),
		x.Child(x.Path(x.D("M21 12h1"))),
		x.Child(x.Path(x.D("M18.5 4.5 18 5"))),
		x.Child(x.Path(x.D("M2 12h1"))),
		x.Child(x.Path(x.D("M12 2v1"))),
		x.Child(x.Path(x.D("m4.929 4.929.707.707"))),
		x.Child(x.Path(x.D("M12 12v6"))),
	)
	return x.Svg(svgArgs...)
}
