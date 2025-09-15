package lucide

import x "github.com/plainkit/html"

// CloudFog creates a Cloud Fog Lucide icon.
func CloudFog(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cloud-fog", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242"))),
		x.Child(x.Path(x.D("M16 17H7"))),
		x.Child(x.Path(x.D("M17 21H9"))),
	)
	return x.Svg(svgArgs...)
}
