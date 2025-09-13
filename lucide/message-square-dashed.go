package lucide

import x "github.com/bloxui/blox"

// MessageSquareDashed creates a Message Square Dashed Lucide icon.
func MessageSquareDashed(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-message-square-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 19h.01"))),
		x.Child(x.Path(x.D("M12 3h.01"))),
		x.Child(x.Path(x.D("M16 19h.01"))),
		x.Child(x.Path(x.D("M16 3h.01"))),
		x.Child(x.Path(x.D("M2 13h.01"))),
		x.Child(x.Path(x.D("M2 17v4.286a.71.71 0 0 0 1.212.502l2.202-2.202A2 2 0 0 1 6.828 19H8"))),
		x.Child(x.Path(x.D("M2 5a2 2 0 0 1 2-2"))),
		x.Child(x.Path(x.D("M2 9h.01"))),
		x.Child(x.Path(x.D("M20 3a2 2 0 0 1 2 2"))),
		x.Child(x.Path(x.D("M22 13h.01"))),
		x.Child(x.Path(x.D("M22 17a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M22 9h.01"))),
		x.Child(x.Path(x.D("M8 3h.01"))),
	)
	return x.Svg(svgArgs...)
}
