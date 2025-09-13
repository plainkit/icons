package lucide

import x "github.com/bloxui/blox"

// Award creates an Award Lucide icon.
func Award(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-award", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m15.477 12.89 1.515 8.526a.5.5 0 0 1-.81.47l-3.58-2.687a1 1 0 0 0-1.197 0l-3.586 2.686a.5.5 0 0 1-.81-.469l1.514-8.526"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("8"), x.R("6"))),
	)
	return x.Svg(svgArgs...)
}
