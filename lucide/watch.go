package lucide

import x "github.com/bloxui/blox"

// Watch creates a Watch Lucide icon.
func Watch(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-watch", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 10v2.2l1.6 1"))),
		x.Child(x.Path(x.D("m16.13 7.66-.81-4.05a2 2 0 0 0-2-1.61h-2.68a2 2 0 0 0-2 1.61l-.78 4.05"))),
		x.Child(x.Path(x.D("m7.88 16.36.8 4a2 2 0 0 0 2 1.61h2.72a2 2 0 0 0 2-1.61l.81-4.05"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("6"))),
	)
	return x.Svg(svgArgs...)
}
