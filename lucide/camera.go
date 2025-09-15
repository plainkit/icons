package lucide

import x "github.com/plainkit/html"

// Camera creates a Camera Lucide icon.
func Camera(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-camera", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.997 4a2 2 0 0 1 1.76 1.05l.486.9A2 2 0 0 0 18.003 7H20a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1.997a2 2 0 0 0 1.759-1.048l.489-.904A2 2 0 0 1 10.004 4z"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("13"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
