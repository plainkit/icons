package lucide

import x "github.com/bloxui/blox"

// FileKey2 creates a File Key 2 Lucide icon.
func FileKey2(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-key-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v6"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Circle(x.Cx("4"), x.Cy("16"), x.R("2"))),
		x.Child(x.Path(x.D("m10 10-4.5 4.5"))),
		x.Child(x.Path(x.D("m9 11 1 1"))),
	)
	return x.Svg(svgArgs...)
}
