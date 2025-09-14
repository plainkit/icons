package lucide

import x "github.com/bloxui/blox"

// SwitchCamera creates a Switch Camera Lucide icon.
func SwitchCamera(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-switch-camera", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 19H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h5"))),
		x.Child(x.Path(x.D("M13 5h7a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-5"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("3"))),
		x.Child(x.Path(x.D("m18 22-3-3 3-3"))),
		x.Child(x.Path(x.D("m6 2 3 3-3 3"))),
	)
	return x.Svg(svgArgs...)
}
