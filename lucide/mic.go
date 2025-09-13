package lucide

import x "github.com/bloxui/blox"

// Mic creates a Mic Lucide icon.
func Mic(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-mic", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 19v3"))),
		x.Child(x.Path(x.D("M19 10v2a7 7 0 0 1-14 0v-2"))),
		x.Child(x.Rect(x.X("9"), x.Y("2"), x.RectWidth("6"), x.RectHeight("13"), x.Rx("3"))),
	)
	return x.Svg(svgArgs...)
}
