package lucide

import x "github.com/bloxui/blox"

// Skull creates a Skull Lucide icon.
func Skull(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-skull", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12.5 17-.5-1-.5 1h1z"))),
		x.Child(x.Path(x.D("M15 22a1 1 0 0 0 1-1v-1a2 2 0 0 0 1.56-3.25 8 8 0 1 0-11.12 0A2 2 0 0 0 8 20v1a1 1 0 0 0 1 1z"))),
		x.Child(x.Circle(x.Cx("15"), x.Cy("12"), x.R("1"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("12"), x.R("1"))),
	)
	return x.Svg(svgArgs...)
}
