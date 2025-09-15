package lucide

import x "github.com/plainkit/html"

// FileType2 creates a File Type 2 Lucide icon.
func FileType2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-type-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M2 13v-1h6v1"))),
		x.Child(x.Path(x.D("M5 12v6"))),
		x.Child(x.Path(x.D("M4 18h2"))),
	)
	return x.Svg(svgArgs...)
}
