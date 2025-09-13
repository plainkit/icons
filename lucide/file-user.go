package lucide

import x "github.com/bloxui/blox"

// FileUser creates a File User Lucide icon.
func FileUser(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-user", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M15 18a3 3 0 1 0-6 0"))),
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7z"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("13"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
