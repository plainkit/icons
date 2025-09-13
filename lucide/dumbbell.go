package lucide

import x "github.com/bloxui/blox"

// Dumbbell creates a Dumbbell Lucide icon.
func Dumbbell(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-dumbbell", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17.596 12.768a2 2 0 1 0 2.829-2.829l-1.768-1.767a2 2 0 0 0 2.828-2.829l-2.828-2.828a2 2 0 0 0-2.829 2.828l-1.767-1.768a2 2 0 1 0-2.829 2.829z"))),
		x.Child(x.Path(x.D("m2.5 21.5 1.4-1.4"))),
		x.Child(x.Path(x.D("m20.1 3.9 1.4-1.4"))),
		x.Child(x.Path(x.D("M5.343 21.485a2 2 0 1 0 2.829-2.828l1.767 1.768a2 2 0 1 0 2.829-2.829l-6.364-6.364a2 2 0 1 0-2.829 2.829l1.768 1.767a2 2 0 0 0-2.828 2.829z"))),
		x.Child(x.Path(x.D("m9.6 14.4 4.8-4.8"))),
	)
	return x.Svg(svgArgs...)
}
