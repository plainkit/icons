package lucide

import x "github.com/bloxui/blox"

// ArrowLeftFromLine creates a Arrow Left From Line Lucide icon.
func ArrowLeftFromLine(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-arrow-left-from-line", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m9 6-6 6 6 6"))),
		x.Child(x.Path(x.D("M3 12h14"))),
		x.Child(x.Path(x.D("M21 19V5"))),
	)
	return x.Svg(svgArgs...)
}
