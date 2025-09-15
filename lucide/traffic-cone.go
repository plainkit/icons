package lucide

import x "github.com/plainkit/html"

// TrafficCone creates a Traffic Cone Lucide icon.
func TrafficCone(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-traffic-cone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16.05 10.966a5 2.5 0 0 1-8.1 0"))),
		x.Child(x.Path(x.D("m16.923 14.049 4.48 2.04a1 1 0 0 1 .001 1.831l-8.574 3.9a2 2 0 0 1-1.66 0l-8.574-3.91a1 1 0 0 1 0-1.83l4.484-2.04"))),
		x.Child(x.Path(x.D("M16.949 14.14a5 2.5 0 1 1-9.9 0L10.063 3.5a2 2 0 0 1 3.874 0z"))),
		x.Child(x.Path(x.D("M9.194 6.57a5 2.5 0 0 0 5.61 0"))),
	)
	return x.Svg(svgArgs...)
}
