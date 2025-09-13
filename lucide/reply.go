package lucide

import x "github.com/bloxui/blox"

// Reply creates a Reply Lucide icon.
func Reply(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-reply", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 18v-2a4 4 0 0 0-4-4H4"))),
		x.Child(x.Path(x.D("m9 17-5-5 5-5"))),
	)
	return x.Svg(svgArgs...)
}
