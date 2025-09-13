package lucide

import x "github.com/bloxui/blox"

// FileX creates a File X Lucide icon.
func FileX(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("m14.5 12.5-5 5"))),
		x.Child(x.Path(x.D("m9.5 12.5 5 5"))),
	)
	return x.Svg(svgArgs...)
}
