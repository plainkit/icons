package lucide

import x "github.com/plainkit/html"

// CloudDrizzle creates a Cloud Drizzle Lucide icon.
func CloudDrizzle(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-drizzle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"))),
		x.Child(x.Path(x.D("M8 19v1"))),
		x.Child(x.Path(x.D("M8 14v1"))),
		x.Child(x.Path(x.D("M16 19v1"))),
		x.Child(x.Path(x.D("M16 14v1"))),
		x.Child(x.Path(x.D("M12 21v1"))),
		x.Child(x.Path(x.D("M12 16v1"))),
	)
	return x.Svg(svgArgs...)
}
