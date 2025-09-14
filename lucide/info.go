package lucide

import x "github.com/bloxui/blox"

// Info creates a Info Lucide icon.
func Info(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-info", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("10"))),
		x.Child(x.Path(x.D("M12 16v-4"))),
		x.Child(x.Path(x.D("M12 8h.01"))),
	)
	return x.Svg(svgArgs...)
}
