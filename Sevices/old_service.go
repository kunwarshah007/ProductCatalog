package Sevices
//
//import (
//	"Interface/dal"
//	"Interface/models"
//	"fmt"
//)
//
//
//type ProductServiceInterface interface {
//
//
//	//Product Sevices Function
//	AddNewProduct		(product models.Product,id int) (err int)
//	GetAllProduct		() []models.Product
//	GetAvailableProduct	() []models.Product
//	GetNewProductById	(id int) 	(p models.Product, err int)
//	BuyProduct		(quantity models.Product,id int) int
//
//
//}
//
//
//
//
//type ProductService struct{
//
//	prod 	  *models.Items
//	trans     *models.Transactions
//
//}
//
//
//
//func NewProductServices(product *models.Items, trans *models.Transactions) *ProductService {
//
//	return &ProductService{product, trans}
//}
//
//
//
//func (s ProductService) AddNewProduct(product models.Product,id int) (err int) {
//
//
//
//	//d := (models.db).Where("ID=?", Id).Find(&getProduct)
//	//_,er:= models.GetProductById(product.Id)
//	//if er!= nil{
//	//	product.CreateProduct()
//	//}
//	//return 1
//	if _, ok := s.prod.Products[id]; ok {
//		if product.ProductName != (s.prod.Products[id]).ProductName || product.Description != (s.prod.Products[id]).Description{
//			//utility.Response(w,400,"Product Not available")
//			return 0
//		}
//		product.Quantity+=(s.prod.Products[id]).Quantity
//		s.prod.Products[id]=product
//		return 1
//	}
//
//	s.prod.Products[id]=product
//	return 2
//}
//
//
//
//func (s ProductService) GetNewProductById(id int) (p models.Product, err int) {
//	if _,ok := s.prod.Products[id] ; !ok{
//		return models.Product{} , 0
//	}
//	return s.prod.Products[id], 1
//}
//
//
//
//
//func (s ProductService) GetAllProduct() []models.Product {
//	TempProduct :=make([]models.Product,0)
//	for _,val := range s.prod.Products{
//		TempProduct =append(TempProduct,val)
//	}
//	fmt.Println(TempProduct)
//	return TempProduct
//}
//func (s ProductService) GetAllTransaction() []models.Bills {
//	TempTransactions :=make([]models.Bills,0)
//	for _,val := range s.trans.Transaction{
//		TempTransactions =append(TempTransactions,val)
//	}
//	fmt.Println(TempTransactions)
//	return TempTransactions
//}
//
//
//func (s ProductService) GetAvailableProduct() []models.Product {
//	TempProduct:=make([]models.Product,0)
//	for _,item :=range s.prod.Products{
//		if item.Quantity>0{
//			TempProduct = append(TempProduct, item)
//		}
//	}
//	return TempProduct
//}
//
//func (s ProductService) BuyProduct(quantity models.Product,id int) int {
//	if _, ok := s.prod.Products[id]; !ok {
//		return 1
//	} else if s.prod.Products[id].Quantity<quantity.Quantity {
//		return 2
//	}else {
//		ProductNewState := s.prod.Products[id]
//		ProductNewState.Quantity-=quantity.Quantity
//		s.prod.Products[id]= ProductNewState
//		var newT models.Bills
//		newT.Id= ProductNewState.Id
//		newT.ProductName=ProductNewState.ProductName
//		newT.Price=ProductNewState.Price
//		newT.Quantity=quantity.Quantity
//		newT.Total= ProductNewState.Price * quantity.Quantity
//
//		//s.trans.Transaction[a] = newT
//		//len(s.trans.Transaction[])
//		return 3
//	}
//}