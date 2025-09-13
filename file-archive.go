package lucide

import x "github.com/bloxui/blox"

// FileArchive creates a File Archive Lucide icon.
func FileArchive(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-archive", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 12v-1"))),
		x.Child(x.Path(x.D("M10 18v-2"))),
		x.Child(x.Path(x.D("M10 7V6"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M15.5 22H18a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v16a2 2 0 0 0 .274 1.01"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("20"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}