package text2img

import (
	"image/color"
	"math/rand"
)

// Color contains a good conbination of backgroundColor and textColor
type Color struct {
	BackgroundColor color.RGBA
	TextColor       color.RGBA
}

var colors []Color

func init() {
	g333 := must(Hex("#333"))
	fff := must(Hex("#fff"))
	colors = []Color{
		{must(Hex("#003d47")), fff},
		{must(Hex("#128277")), fff},
		{must(Hex("#d24136")), fff},
		{must(Hex("#eb8a3e")), fff},
		{must(Hex("#ebb582")), fff},
		{must(Hex("#785a46")), fff},
		{must(Hex("#bc6d4f")), fff},
		{must(Hex("#1e1f26")), fff},
		{must(Hex("#283655")), fff},
		{must(Hex("#4d648d")), fff},
		{must(Hex("#265c00")), fff},
		{must(Hex("#faaf08")), fff},
		{must(Hex("#fa812f")), fff},
		{must(Hex("#fa4032")), fff},
		{must(Hex("#6c5f5b")), fff},
		{must(Hex("#cdab81")), fff},
		{must(Hex("#4f4a45")), fff},
		{must(Hex("#04202c")), fff},
		{must(Hex("#304040")), fff},
		{must(Hex("#5b7065")), fff},
		{must(Hex("#1e0000")), fff},
		{must(Hex("#500805")), fff},
		{must(Hex("#9d331f")), fff},
		{must(Hex("#68a225")), fff},
		{must(Hex("#fdffff")), g333},
		{must(Hex("#2c4a52")), fff},
		{must(Hex("#537072")), fff},
		{must(Hex("#8e9b97")), fff},
		{must(Hex("#f4ebdb")), g333},
		{must(Hex("#d8412f")), fff},
		{must(Hex("#fe7a47")), fff},
		{must(Hex("#fcfdfe")), g333},
		{must(Hex("#867666")), fff},
		{must(Hex("#e1b80d")), fff},
		{must(Hex("#003b46")), fff},
		{must(Hex("#07575b")), fff},
		{must(Hex("#66a5ad")), fff},
		{must(Hex("#af6c59")), fff},
		{must(Hex("#e68f71")), fff},
		{must(Hex("#021c1e")), fff},
		{must(Hex("#004445")), fff},
		{must(Hex("#2c7873")), fff},
		{must(Hex("#6fb98f")), fff},
		{must(Hex("#434343")), fff},
		{must(Hex("#767676")), fff},
		{must(Hex("#c16707")), fff},
		{must(Hex("#f08d16")), fff},
		{must(Hex("#77262a")), fff},
		{must(Hex("#9e2d29")), fff},
		{must(Hex("#c35d44")), fff},
		{must(Hex("#202d35")), fff},
		{must(Hex("#0e3c54")), fff},
		{must(Hex("#2a677c")), fff},
		{must(Hex("#4f3538")), fff},
		{must(Hex("#66443b")), fff},
		{must(Hex("#c29f83")), fff},
		{must(Hex("#210e3b")), fff},
		{must(Hex("#4b194c")), fff},
		{must(Hex("#872b76")), fff},
	}
}

// PickColor picks a color
func PickColor() Color {
	return colors[rand.Intn(len(colors))]
}
