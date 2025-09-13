package lucide

import x "github.com/bloxui/blox"

// BadgeIndianRupee creates a Badge Indian Rupee Lucide icon.
func BadgeIndianRupee(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-badge-indian-rupee", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z"))),
		x.Child(x.Path(x.D("M8 8h8"))),
		x.Child(x.Path(x.D("M8 12h8"))),
		x.Child(x.Path(x.D("m13 17-5-1h1a4 4 0 0 0 0-8"))),
	)
	return x.Svg(svgArgs...)
}
