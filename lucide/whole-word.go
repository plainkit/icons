package lucide

import x "github.com/plainkit/html"

// WholeWord creates a Whole Word Lucide icon.
func WholeWord(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-whole-word", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("7"), x.Cy("12"), x.R("3"))),
		x.Child(x.Path(x.D("M10 9v6"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("12"), x.R("3"))),
		x.Child(x.Path(x.D("M14 7v8"))),
		x.Child(x.Path(x.D("M22 17v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"))),
	)
	return x.Svg(svgArgs...)
}
