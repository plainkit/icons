package lucide

import x "github.com/plainkit/html"

// Joystick creates a Joystick Lucide icon.
func Joystick(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-joystick", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-2Z"))),
		x.Child(x.Path(x.D("M6 15v-2"))),
		x.Child(x.Path(x.D("M12 15V9"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("6"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
