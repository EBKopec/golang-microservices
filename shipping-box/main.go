package shipping_box

type Product struct {
	Name   string
	Length int
	Width  int
	Height int
	Weight int
}

type Box struct {
	Length int
	Width  int
	Height int
	Weight int
}

func getBestBox(availableBoxes []Box, products []Product) Box {
	//Todo: Complete!
	return Box{}
}
