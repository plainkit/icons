package lucide

import x "github.com/bloxui/blox"

// ArrowUpZA creates a Arrow Up Z A Lucide icon.
func ArrowUpZA(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up-z-a", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m3 8 4-4 4 4"))),
		x.Child(x.Path(x.D("M7 4v16"))),
		x.Child(x.Path(x.D("M15 4h5l-5 6h5"))),
		x.Child(x.Path(x.D("M15 20v-3.5a2.5 2.5 0 0 1 5 0V20"))),
		x.Child(x.Path(x.D("M20 18h-5"))),
	)
	return x.Svg(svgArgs...)
}
