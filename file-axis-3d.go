package lucide

import x "github.com/bloxui/blox"

// FileAxis3d creates a File Axis 3d Lucide icon.
func FileAxis3d(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-axis-3d", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("m8 18 4-4"))),
		x.Child(x.Path(x.D("M8 10v8h8"))),
	)
	return x.Svg(svgArgs...)
}