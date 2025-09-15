package lucide

import x "github.com/plainkit/blox"

// HandPlatter creates a Hand Platter Lucide icon.
func HandPlatter(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-hand-platter", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3V2"))),
		x.Child(x.Path(x.D("m15.4 17.4 3.2-2.8a2 2 0 1 1 2.8 2.9l-3.6 3.3c-.7.8-1.7 1.2-2.8 1.2h-4c-1.1 0-2.1-.4-2.8-1.2l-1.302-1.464A1 1 0 0 0 6.151 19H5"))),
		x.Child(x.Path(x.D("M2 14h12a2 2 0 0 1 0 4h-2"))),
		x.Child(x.Path(x.D("M4 10h16"))),
		x.Child(x.Path(x.D("M5 10a7 7 0 0 1 14 0"))),
		x.Child(x.Path(x.D("M5 14v6a1 1 0 0 1-1 1H2"))),
	)
	return x.Svg(svgArgs...)
}
