package lucide

import x "github.com/plainkit/html"

// Flower2 creates a Flower 2 Lucide icon.
func Flower2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-flower-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 5a3 3 0 1 1 3 3m-3-3a3 3 0 1 0-3 3m3-3v1M9 8a3 3 0 1 0 3 3M9 8h1m5 0a3 3 0 1 1-3 3m3-3h-1m-2 3v-1"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("8"), x.R("2"))),
		x.Child(x.Path(x.D("M12 10v12"))),
		x.Child(x.Path(x.D("M12 22c4.2 0 7-1.667 7-5-4.2 0-7 1.667-7 5Z"))),
		x.Child(x.Path(x.D("M12 22c-4.2 0-7-1.667-7-5 4.2 0 7 1.667 7 5Z"))),
	)
	return x.Svg(svgArgs...)
}
