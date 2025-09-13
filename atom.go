package lucide

import x "github.com/bloxui/blox"

// Atom creates a Atom Lucide icon.
func Atom(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-atom", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("1"))),
		x.Child(x.Path(x.D("M20.2 20.2c2.04-2.03.02-7.36-4.5-11.9-4.54-4.52-9.87-6.54-11.9-4.5-2.04 2.03-.02 7.36 4.5 11.9 4.54 4.52 9.87 6.54 11.9 4.5Z"))),
		x.Child(x.Path(x.D("M15.7 15.7c4.52-4.54 6.54-9.87 4.5-11.9-2.03-2.04-7.36-.02-11.9 4.5-4.52 4.54-6.54 9.87-4.5 11.9 2.03 2.04 7.36.02 11.9-4.5Z"))),
	)
	return x.Svg(svgArgs...)
}
