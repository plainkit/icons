package lucide

import x "github.com/plainkit/html"

// VolumeOff creates a Volume Off Lucide icon.
func VolumeOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-volume-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 9a5 5 0 0 1 .95 2.293"))),
		x.Child(x.Path(x.D("M19.364 5.636a9 9 0 0 1 1.889 9.96"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("m7 7-.587.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298V11"))),
		x.Child(x.Path(x.D("M9.828 4.172A.686.686 0 0 1 11 4.657v.686"))),
	)
	return x.Svg(svgArgs...)
}
