package lucide

import x "github.com/bloxui/blox"

// MessageSquareShare creates a Message Square Share Lucide icon.
func MessageSquareShare(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-message-square-share", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 3H4a2 2 0 0 0-2 2v16.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H20a2 2 0 0 0 2-2v-4"))),
		x.Child(x.Path(x.D("M16 3h6v6"))),
		x.Child(x.Path(x.D("m16 9 6-6"))),
	)
	return x.Svg(svgArgs...)
}
