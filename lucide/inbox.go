package lucide

import x "github.com/bloxui/blox"

// Inbox creates a Inbox Lucide icon.
func Inbox(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-inbox", args)
	svgArgs = append(svgArgs,
		x.Child(x.Polyline(x.Points("22 12 16 12 14 15 10 15 8 12 2 12"))),
		x.Child(x.Path(x.D("M5.45 5.11 2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z"))),
	)
	return x.Svg(svgArgs...)
}
