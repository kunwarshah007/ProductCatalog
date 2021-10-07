package test

import (
	"Interface/Handler"
	"Interface/Models"
	mock_services "Interface/mock/Sevices"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateProduct(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	InputProduct:=Models.Product{Name:"wickets",Description :"wood",Price: 256,Quantity: 3}
	//OutputProduct:=Models.Product{Id:1,ProductName:"wickets",Description :"wood",Price: 256,Quantity: 3}
	BodyProduct := []byte(`{"name":"wickets","description":"wood","price":256,"quantity":3}`)

	req, _ := http.NewRequest("POST", "/products/insert", bytes.NewBuffer(BodyProduct))
	MockServices := mock_services.NewMockProductServiceInterface(ctrl)
	MockServices.EXPECT().AddNewProduct(InputProduct).Return(1,0)
	ProductController:=Handler.Initialise(MockServices)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/products/insert",ProductController.CreateProduct ).Methods("POST")
	r.ServeHTTP(rr, req)

	StringResponse:=`{"Product created with ID":1}`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 200,rr.Code)
}

func TestCreateProduct2(t *testing.T){
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish()

	InputProduct:=Models.Product{Name:"wickets",Description :"wood",Price: 256,Quantity: 3}
	//OutputProduct:=Models.Product{Id:1,ProductName:"wickets",Description :"wood",Price: 256,Quantity: 3}
	BodyProduct := []byte(`{"name":"wickets","description":"wood","price":256,"quantity":3}`)

	req, _ := http.NewRequest("POST", "/products/insert", bytes.NewBuffer(BodyProduct))
	MockServices := mock_services.NewMockProductServiceInterface(ctrl)
	MockServices.EXPECT().AddNewProduct(InputProduct).Return(1,1)
	ProductController:=Handler.Initialise(MockServices)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()

	r.HandleFunc("/products/insert",ProductController.CreateProduct ).Methods("POST")
	r.ServeHTTP(rr, req)

	StringResponse:=`{"Result-\u003e":"Product with given id already present in the database!!!"}`
	str:=strings.TrimSpace(rr.Body.String())

	assert.Equal(t, StringResponse,str)
	assert.Equal(t, 400,rr.Code)
}

