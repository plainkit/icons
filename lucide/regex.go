package lucide

import x "github.com/bloxui/blox"

// Regex creates a Regex Lucide icon.
func Regex(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-regex", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 3v10"))),
		x.Child(x.Path(x.D("m12.67 5.5 8.66 5"))),
		x.Child(x.Path(x.D("m12.67 10.5 8.66-5"))),
		x.Child(x.Path(x.D("M9 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2z"))),
	)
	return x.Svg(svgArgs...)
}
