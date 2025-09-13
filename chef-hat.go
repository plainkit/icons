package lucide

import x "github.com/bloxui/blox"

// ChefHat creates a Chef Hat Lucide icon.
func ChefHat(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-chef-hat", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M17 21a1 1 0 0 0 1-1v-5.35c0-.457.316-.844.727-1.041a4 4 0 0 0-2.134-7.589 5 5 0 0 0-9.186 0 4 4 0 0 0-2.134 7.588c.411.198.727.585.727 1.041V20a1 1 0 0 0 1 1Z"))),
		x.Child(x.Path(x.D("M6 17h12"))),
	)
	return x.Svg(svgArgs...)
}
