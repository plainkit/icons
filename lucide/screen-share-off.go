package lucide

import x "github.com/plainkit/blox"

// ScreenShareOff creates a Screen Share Off Lucide icon.
func ScreenShareOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-screen-share-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3"))),
		x.Child(x.Path(x.D("M8 21h8"))),
		x.Child(x.Path(x.D("M12 17v4"))),
		x.Child(x.Path(x.D("m22 3-5 5"))),
		x.Child(x.Path(x.D("m17 3 5 5"))),
	)
	return x.Svg(svgArgs...)
}
