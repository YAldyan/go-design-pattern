package visitor

type Visitable interface {
	Accept(Visitor)
}

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}
func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}
func (p *Product) GetName() string {
	return p.Name
}

//Generic Product Created
type Rice struct {
	Product
}

type Pasta struct {
	Product
}

/*
	Ketika hanya perlu untuk override GetPrice
	dikarenakan adanya tambahan biaya.
*/
type Fridge struct {
	Product
}

func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

/*
	Tidak cukup hanya tambah GetPrice
	dikarenakan kita juga perlu menjadikan
	Fridge anak Visitable Interface
*/
func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}

/*Interface Objek yang mengunjungi*/
type Visitor interface {
	Visit(ProductInfoRetriever)
}

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.Names = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}
