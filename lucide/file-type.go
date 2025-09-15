package lucide

import x "github.com/plainkit/html"

// FileType creates a File Type Lucide icon.
func FileType(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-type", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M9 13v-1h6v1"))),
		x.Child(x.Path(x.D("M12 12v6"))),
		x.Child(x.Path(x.D("M11 18h2"))),
	)
	return x.Svg(svgArgs...)
}
