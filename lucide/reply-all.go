package lucide

import x "github.com/bloxui/blox"

// ReplyAll creates a Reply All Lucide icon.
func ReplyAll(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-reply-all", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12 17-5-5 5-5"))),
		x.Child(x.Path(x.D("M22 18v-2a4 4 0 0 0-4-4H7"))),
		x.Child(x.Path(x.D("m7 17-5-5 5-5"))),
	)
	return x.Svg(svgArgs...)
}
