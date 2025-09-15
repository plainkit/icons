package lucide

import x "github.com/plainkit/html"

// CodeXml creates a Code Xml Lucide icon.
func CodeXml(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-code-xml", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m18 16 4-4-4-4"))),
		x.Child(x.Path(x.D("m6 8-4 4 4 4"))),
		x.Child(x.Path(x.D("m14.5 4-5 16"))),
	)
	return x.Svg(svgArgs...)
}
