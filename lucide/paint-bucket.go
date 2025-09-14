package lucide

import x "github.com/bloxui/blox"

// PaintBucket creates a Paint Bucket Lucide icon.
func PaintBucket(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-paint-bucket", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m19 11-8-8-8.6 8.6a2 2 0 0 0 0 2.8l5.2 5.2c.8.8 2 .8 2.8 0L19 11Z"))),
		x.Child(x.Path(x.D("m5 2 5 5"))),
		x.Child(x.Path(x.D("M2 13h15"))),
		x.Child(x.Path(x.D("M22 20a2 2 0 1 1-4 0c0-1.6 1.7-2.4 2-4 .3 1.6 2 2.4 2 4Z"))),
	)
	return x.Svg(svgArgs...)
}
