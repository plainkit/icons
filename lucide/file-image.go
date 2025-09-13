package lucide

import x "github.com/bloxui/blox"

// FileImage creates a File Image Lucide icon.
func FileImage(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-file-image", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("12"), x.R("2"))),
		x.Child(x.Path(x.D("m20 17-1.296-1.296a2.41 2.41 0 0 0-3.408 0L9 22"))),
	)
	return x.Svg(svgArgs...)
}
