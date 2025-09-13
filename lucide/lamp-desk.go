package lucide

import x "github.com/bloxui/blox"

// LampDesk creates a Lamp Desk Lucide icon.
func LampDesk(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-lamp-desk", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.293 2.293a1 1 0 0 1 1.414 0l2.5 2.5 5.994 1.227a1 1 0 0 1 .506 1.687l-7 7a1 1 0 0 1-1.687-.506l-1.227-5.994-2.5-2.5a1 1 0 0 1 0-1.414z"))),
		x.Child(x.Path(x.D("m14.207 4.793-3.414 3.414"))),
		x.Child(x.Path(x.D("M3 20a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"))),
		x.Child(x.Path(x.D("m9.086 6.5-4.793 4.793a1 1 0 0 0-.18 1.17L7 18"))),
	)
	return x.Svg(svgArgs...)
}
