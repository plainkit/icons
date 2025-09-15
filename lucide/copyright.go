package lucide

import x "github.com/plainkit/html"

// Copyright creates a Copyright Lucide icon.
func Copyright(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-copyright", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M14.83 14.83a4 4 0 1 1 0-5.66"))),
	)
	return x.Svg(svgArgs...)
}
