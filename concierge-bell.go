package lucide

import x "github.com/bloxui/blox"

// ConciergeBell creates a Concierge Bell Lucide icon.
func ConciergeBell(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-concierge-bell", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 20a1 1 0 0 1-1-1v-1a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1Z"))),
		x.Child(x.Path(x.D("M20 16a8 8 0 1 0-16 0"))),
		x.Child(x.Path(x.D("M12 4v4"))),
		x.Child(x.Path(x.D("M10 4h4"))),
	)
	return x.Svg(svgArgs...)
}