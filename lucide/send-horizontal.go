package lucide

import x "github.com/bloxui/blox"

// SendHorizontal creates a Send Horizontal Lucide icon.
func SendHorizontal(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-send-horizontal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.714 3.048a.498.498 0 0 0-.683.627l2.843 7.627a2 2 0 0 1 0 1.396l-2.842 7.627a.498.498 0 0 0 .682.627l18-8.5a.5.5 0 0 0 0-.904z"))),
		x.Child(x.Path(x.D("M6 12h16"))),
	)
	return x.Svg(svgArgs...)
}
