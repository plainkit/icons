package lucide

import x "github.com/bloxui/blox"

// FilePlus2 creates a File Plus 2 Lucide icon.
func FilePlus2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-plus-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M3 15h6"))),
		x.Child(x.Path(x.D("M6 12v6"))),
	)
	return x.Svg(svgArgs...)
}
