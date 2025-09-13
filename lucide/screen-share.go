package lucide

import x "github.com/bloxui/blox"

// ScreenShare creates a Screen Share Lucide icon.
func ScreenShare(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-screen-share", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3"))),
		x.Child(x.Path(x.D("M8 21h8"))),
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("m17 8 5-5"))),
		x.Child(x.Path(x.D("M17 3h5v5"))),
	)
	return x.Svg(svgArgs...)
}
