package lucide

import x "github.com/bloxui/blox"

// DraftingCompass creates a Drafting Compass Lucide icon.
func DraftingCompass(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-drafting-compass", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m12.99 6.74 1.93 3.44"))),
		x.Child(x.Path(x.D("M19.136 12a10 10 0 0 1-14.271 0"))),
		x.Child(x.Path(x.D("m21 21-2.16-3.84"))),
		x.Child(x.Path(x.D("m3 21 8.02-14.26"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("5"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
