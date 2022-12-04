package application_test

import (
	"github.com/LuizFernandesOliveira/fullcycle-3-hexagonal-architecture-go/application"
	mock_application "github.com/LuizFernandesOliveira/fullcycle-3-hexagonal-architecture-go/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Get(gomock.All()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	result, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Save(gomock.All()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	result, err := service.Create("Product 1", 10.0)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Save(gomock.All()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
