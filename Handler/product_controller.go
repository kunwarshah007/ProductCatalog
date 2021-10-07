package Handler

import (
	"Interface/Handler/ErrorHandler"
	"Interface/Models"
	"Interface/Sevices"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)
type ControllerInterface interface {
	GetProductById(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetAllProduct(w http.ResponseWriter, r *http.Request)
	BuyProduct(w http.ResponseWriter, r *http.Request)
	GetAllTransactions(w http.ResponseWriter, r *http.Request)
	GetTop5Products(w http.ResponseWriter, r *http.Request)
}

type ProductController struct{
	 ps  Sevices.ProductServiceInterface
}


func Initialise(p Sevices.ProductServiceInterface) ControllerInterface {
	return &ProductController{p}
}

type top struct{
	Id 			int 	`json:"id"`
	Name 		string 	`json:"name"`
	Quantity 	int 	`json:"quantity"`
}


func (pc ProductController) GetProductById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var id int
		var e error
	    id, e = strconv.Atoi(params["id"])
		if e != nil{
			ErrorHandler.Response(w,400,"Product Id should be Integer","501")
			return
		}

		IdProduct ,err :=pc.ps.GetNewProductById(id)
		if err==1 {
			ErrorHandler.Response(w,400,"Product With Given Id is not present in the store","502")
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(IdProduct)
}

func (pc ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var product Models.Product
		e := json.NewDecoder(r.Body).Decode(&product)
		if e!= nil{
			ErrorHandler.Response(w,400,"Error occur while Decoding: Id , Price ,Quantity should be int ","503")
			return
		}
	    i ,err :=pc.ps.AddNewProduct(product)

		if err==1 {
			ErrorHandler.Response(w,400,"Product with given id already present in the database!!!","504")
			return
		}
		//msg:="Product with id:"+(string)(i)+"is created  Successfully and inserted into the Store"
		//err=json.NewEncoder(w).Encode((string)(msg))
		ErrorHandler.Response1(w,200,"Product Successfully inserted into the Store",i,product.Name)

}


func (pc ProductController) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p:= pc.ps.GetAllProduct()
	err:=json.NewEncoder(w).Encode(p)
	if err!= nil{
		ErrorHandler.Response(w,400,"Error occur while Encoding","505")
		return
	}
	w.WriteHeader(200)
}


func (pc ProductController) BuyProduct(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var value Models.Product
	e := json.NewDecoder(r.Body).Decode(&value)
	if e!= nil{
		ErrorHandler.Response(w,400,"Error occur while Decoding: Id , Price ,Quantity should be int ","503")
		return
	}
	var id int
	//var e error
	id, e = strconv.Atoi(params["id"])
	if e != nil{
		ErrorHandler.Response(w,400,"Product Id should be Integer","501")
		return
	}
	err,total,name:=pc.ps.BuyProduct(value,id)
	if err== 1{
		ErrorHandler.Response(w,400,"Product With given Id is not available","506")
		return
	}else if err == 2 {
		ErrorHandler.Response(w,400,"sufficient quantity not available","507")
		return
	}else{
		ErrorHandler.Response2(w,200,"Product purchase successful",id,name,total)
		return
	}
}



//func (pc ProductController) BuyProductMany(w http.ResponseWriter, r *http.Request)  {
//	w.Header().Set("Content-Type", "application/json")
//	//params := mux.Vars(r)
//	var value []Models.Product
//	e := json.NewDecoder(r.Body).Decode(&value)
//	if e!= nil{
//		ErrorHandler.Response(w,400,"Error occur while Decoding: Id , Price ,Quantity should be int ","503")
//		return
//	}
//	var id int
//	//var e error
//	id, e = strconv.Atoi(params["id"])
//	if e != nil{
//		ErrorHandler.Response(w,400,"Product Id should be Integer","501")
//		return
//	}
//	err,total,name:=pc.ps.BuyProduct(value,id)
//	if err== 1{
//		ErrorHandler.Response(w,400,"Product With given Id is not available","506")
//		return
//	}else if err == 2 {
//		ErrorHandler.Response(w,400,"sufficient quantity not available","507")
//		return
//	}else{
//		ErrorHandler.Response2(w,200,"Product purchase successful",id,name,total)
//		return
//	}
//}



func (pc ProductController) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t:= pc.ps.GetAllTransaction()
	err:=json.NewEncoder(w).Encode(t)
	if err!= nil{
		ErrorHandler.Response(w,400,"Error occur while Encoding","505")
		return
	}
	w.WriteHeader(200)
}

func (pc ProductController) GetTop5Products(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	t:= pc.ps.GetTop5Products()
	var top5 []top
	for i,_ :=range t{
		var temp top
		temp.Id=t[i].ProductId
		p,_:=pc.ps.GetNewProductById(t[i].ProductId)
		temp.Name=p.Name
		temp.Quantity=t[i].Quantity
		top5=append(top5,temp)
	}
	err:=json.NewEncoder(w).Encode(top5)
	if err!= nil{
		ErrorHandler.Response(w,400,"Error occur while Encoding","505")
		return
	}
	w.WriteHeader(200)

}
