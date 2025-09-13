package lucide

import x "github.com/bloxui/blox"

// Ampersands creates a Ampersands Lucide icon.
func Ampersands(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-ampersands", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 17c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6 0 1.7 1.3 3 3 3 2.8 0 5-2.2 5-5"))),
		x.Child(x.Path(x.D("M22 17c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6 0 1.7 1.3 3 3 3 2.8 0 5-2.2 5-5"))),
	)
	return x.Svg(svgArgs...)
}
