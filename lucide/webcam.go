package lucide

import x "github.com/bloxui/blox"

// Webcam creates a Webcam Lucide icon.
func Webcam(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-webcam", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("8"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("3"))),
		x.Child(x.Path(x.D("M7 22h10"))),
		x.Child(x.Path(x.D("M12 22v-4"))),
	)
	return x.Svg(svgArgs...)
}
