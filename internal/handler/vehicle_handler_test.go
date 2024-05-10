package handler

import (
	"app/internal"
	"app/internal/service"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
func TestHandlerVehicle_FindByColorAndYear(t *testing.T) {

	type arrange struct {
		svMock func() *service.MockVehicle
	}
	type input struct {
		request  func() *http.Request
		response *httptest.ResponseRecorder
	}
	type output struct {
		statusCode int
		body       string
		headers    http.Header
	}
	type testCase struct {
		name    string
		arrange arrange
		input   input
		output  output
	}

	// Case 1: Valid request - Find by color and year
	testCases := []testCase{
		{
			name: "Valid request - Find by color and year",
			arrange: arrange{
				svMock: func() *service.MockVehicle {
					svMock := service.NewMocksVehicle()
					vehicles := map[int]internal.Vehicle{
						1: {
							Id: 1,
							VehicleAttributes: internal.VehicleAttributes{
								Brand:           "Mustang",
								Model:           "model",
								Registration:    "registration",
								Color:           "red",
								FabricationYear: 2010,
								Capacity:        5,
								MaxSpeed:        100,
								FuelType:        "fuelType",
								Transmission:    "transmission",
								Weight:          1000,
								Dimensions: internal.Dimensions{
									Height: 1,
									Length: 1,
									Width:  1,
								},
							},
						},
						2: {
							Id: 2,
							VehicleAttributes: internal.VehicleAttributes{
								Brand:           "Mustang",
								Model:           "model2",
								Registration:    "registration2",
								Color:           "red",
								FabricationYear: 2010,
								Capacity:        6,
								MaxSpeed:        200,
								FuelType:        "fuelType2",
								Transmission:    "transmission2",
								Weight:          2000,
								Dimensions: internal.Dimensions{
									Height: 2,
									Length: 2,
									Width:  2,
								},
							},
						},
					}
					svMock.On("FindByColorAndYear", "red", 2010).Return(vehicles, nil)
					return svMock
				},
			},
			input: input{
				request: func() *http.Request {
					return httptest.NewRequest("GET", "/vehicles/red/2010", nil)
				},
				response: httptest.NewRecorder(),
			},
			output: output{
				statusCode: http.StatusOK,
				body: `{
"message":"vehicles found",
"data":{
	"1":{
		"id":1,
		"brand":"Mustang",
		"model":"model",
		"registration":"registration",
		"color":"red",
		"fabrication_year":2010,
		"capacity":5,
		"max_speed":100,
		"fuel_type":"fuelType",
		"transmission":"transmission",
		"weight":1000,
		"dimensions":{
			"height":1,
			"length":1,
			"width":1
		}
	},
	"2":{
		"id":2,
		"brand":"Mustang",
		"model":"model2",
		"registration":"registration2",
		"color":"red",
		"fabrication_year":2010,
		"capacity":6,
		"max_speed":200,
		"fuel_type":"fuelType2",
		"transmission":"transmission2",
		"weight":2000,
		"dimensions":{
			"height":2,
			"length":2,
			"width":2
		}
	}
}
}`,
				headers: http.Header{
					"Content-Type": []string{"application/json"},
				},
			},
		},
		// Other case
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			svMock := tc.arrange.svMock()
			h := NewHandlerVehicle(svMock)

			// Act
			h.FindByColorAndYear()(tc.input.response, tc.input.request())
			// Assert
			require.Equal(t, tc.output.statusCode, tc.input.response.Code)
			require.Equal(t, tc.output.body, tc.input.response.Body.String())
			require.Equal(t, tc.output.headers, tc.input.response.Header())
		})

	}
}

func TestHandlerVehicle_FindByColorAndYear2(t *testing.T) {
	// Mock de service
	mockService := service.NewMocksVehicle()
	handler := NewHandlerVehicle(mockService)

	// Datos de prueba
	color := "blue"
	year := 2022
	expectedVehicles := map[int]internal.Vehicle{
		1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "Mustang",
				Model:           "model",
				Registration:    "registration",
				Color:           "red",
				FabricationYear: 2010,
				Capacity:        5,
				MaxSpeed:        100,
				FuelType:        "fuelType",
				Transmission:    "transmission",
				Weight:          1000,
				Dimensions: internal.Dimensions{
					Height: 1,
					Length: 1,
					Width:  1,
				},
			},
		},
		2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "Mustang",
				Model:           "model2",
				Registration:    "registration2",
				Color:           "red",
				FabricationYear: 2010,
				Capacity:        6,
				MaxSpeed:        200,
				FuelType:        "fuelType2",
				Transmission:    "transmission2",
				Weight:          2000,
				Dimensions: internal.Dimensions{
					Height: 2,
					Length: 2,
					Width:  2,
				},
			},
		},
	}

	// Configurar expectativas del mock
	mockService.On("FindByColorAndYear", color, year).Return(expectedVehicles, nil)

	// Crear request
	req, err := http.NewRequest("GET", "/vehicles/blue/2022", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Grabar respuesta del handler
	rr := httptest.NewRecorder()
	handler.FindByColorAndYear().ServeHTTP(rr, req)

	// Verificar código de estado
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verificar el cuerpo de la respuesta
	expectedResponse := `{
"message":"vehicles found",
"data":{
	"1":{
		"id":1,
		"brand":"Mustang",
		"model":"model",
		"registration":"registration",
		"color":"red",
		"fabrication_year":2010,
		"capacity":5,
		"max_speed":100,
		"fuel_type":"fuelType",
		"transmission":"transmission",
		"weight":1000,
		"dimensions":{
			"height":1,
			"length":1,
			"width":1
		}
	},
	"2":{
		"id":2,
		"brand":"Mustang",
		"model":"model2",
		"registration":"registration2",
		"color":"red",
		"fabrication_year":2010,
		"capacity":6,
		"max_speed":200,
		"fuel_type":"fuelType2",
		"transmission":"transmission2",
		"weight":2000,
		"dimensions":{
			"height":2,
			"length":2,
			"width":2
		}
	}
}
}`
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Verificar llamada al mock
	mockService.AssertExpectations(t)
}
*/

// Test for AverageMaxSpeedByBrand: Successful request
func TestHandlerVehicle_AverageMaxSpeedByBrand(t *testing.T) {
	// Mock del servicio
	mockService := service.NewMocksVehicle()
	handler := NewHandlerVehicle(mockService)

	// Configurar expectativas del mock
	mockService.On("AverageMaxSpeedByBrand", "Toyota").Return(50.0, nil)

	// Crear solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/vehicles/average_speed/brand/Toyota", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Grabar la respuesta del handler
	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/vehicles/average_speed/brand/{brand}", handler.AverageMaxSpeedByBrand())
	router.ServeHTTP(rr, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verificar el cuerpo de la respuesta
	expectedResponse := `{"data":50,"message":"average max speed found"}`
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Verificar llamada al mock
	mockService.AssertExpectations(t)
}

// Test for AverageMaxSpeedByBrand: Error getting vehicles by brand
func TestHandlerVehicle_AverageMaxSpeedByBrandErrorGettingVehiclesByBrand(t *testing.T) {
	// Mock del servicio
	mockService := service.NewMocksVehicle()
	handler := NewHandlerVehicle(mockService)

	// Configurar expectativas del mock
	mockService.On("AverageMaxSpeedByBrand", "Toyota").Return(0.0, errors.New("error getting vehicles by brand"))

	// Crear solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/vehicles/average_speed/brand/Toyota", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Grabar la respuesta del handler
	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/vehicles/average_speed/brand/{brand}", handler.AverageMaxSpeedByBrand())
	router.ServeHTTP(rr, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Verificar el cuerpo de la respuesta
	expectedResponse := `{"status":"Internal Server Error","message":"internal error"}`
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Verificar llamada al mock
	mockService.AssertExpectations(t)
}

// Test for AverageMaxSpeedByBrand: No vehicles found by brand
func TestHandlerVehicle_AverageMaxSpeedByBrandNoVehiclesFoundByBrand(t *testing.T) {
	// Mock del servicio
	mockService := service.NewMocksVehicle()
	handler := NewHandlerVehicle(mockService)

	// Configurar expectativas del mock
	mockService.On("AverageMaxSpeedByBrand", "Toyota").Return(0.0, internal.ErrServiceNoVehicles)

	// Crear solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/vehicles/average_speed/brand/Toyota", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Grabar la respuesta del handler - Router es apropiado para test de integración, no para unitarios
	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/vehicles/average_speed/brand/{brand}", handler.AverageMaxSpeedByBrand())
	router.ServeHTTP(rr, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusNotFound, rr.Code)
	// Verificar el cuerpo de la respuesta
	expectedResponse := `{"status":"Not Found","message":"vehicles not found"}`
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Verificar llamada al mock
	mockService.AssertExpectations(t)
}

func TestHandlerVehicle_FindByColorAndYearFoundingAll(t *testing.T) {
	// Mock del servicio
	mockService := service.NewMocksVehicle()
	handler := NewHandlerVehicle(mockService)

	// Datos de prueba
	color := "red"
	year := 2010
	expectedVehicles := map[int]internal.Vehicle{
		1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "Mustang",
				Model:           "model",
				Registration:    "registration",
				Color:           "red",
				FabricationYear: 2010,
				Capacity:        5,
				MaxSpeed:        100,
				FuelType:        "fuelType",
				Transmission:    "transmission",
				Weight:          1000,
				Dimensions: internal.Dimensions{
					Height: 1,
					Length: 1,
					Width:  1,
				},
			},
		},
	}

	// Configurar expectativas del mock
	mockService.On("FindByColorAndYear", color, year).Return(expectedVehicles, nil)

	// Crear solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/vehicles/color/red/year/2010", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Grabar la respuesta del handler
	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/vehicles/color/{color}/year/{year}", handler.FindByColorAndYear())
	router.ServeHTTP(rr, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verificar el cuerpo de la respuesta
	expectedResponse := `{"message":"vehicles found","data":{"1":{"Brand":"Mustang","Model":"model","Registration":"registration","Color":"red","FabricationYear":2010,"Capacity":5,"MaxSpeed":100,"FuelType":"fuelType","Transmission":"transmission","Weight":1000,"Height":1,"Id":1, "Length":1,"Width":1}}}`

	assert.JSONEq(t, expectedResponse, rr.Body.String())
	// Verificar llamada al mock
	mockService.AssertExpectations(t)
}

func TestHandlerVehicle_FindByColorAndYearErrorStatusBadRequestInvalidYear(t *testing.T) {
	// Arrange
	// Mock del servicio
	mockService := service.NewMocksVehicle()
	handler := NewHandlerVehicle(mockService)

	// Crear solicitud HTTP simulada con un año inválido
	req, err := http.NewRequest("GET", "/vehicles/color/red/year/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Grabar la respuesta del handler
	rr := httptest.NewRecorder()
	router := chi.NewRouter()

	// Act
	router.Get("/vehicles/color/{color}/year/{year}", handler.FindByColorAndYear())
	router.ServeHTTP(rr, req)

	// Assert
	// Verificar el código de estado
	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// Verificar el cuerpo de la respuesta
	expectedResponse := `{"status":"Bad Request","message":"invalid year"}`
	assert.Equal(t, expectedResponse, rr.Body.String())
}

func TestHandlerVehicle_FindByColorAndYearInvalidColor(t *testing.T) {
	// Mock del servicio
	mockService := service.NewMocksVehicle()
	handler := NewHandlerVehicle(mockService)

	// Configurar expectativas del mock
	mockService.On("FindByColorAndYear", "invalid", 2010).Return(map[int]internal.Vehicle{}, errors.New("invalid color"))

	// Crear solicitud HTTP simulada con un color no válido
	req, err := http.NewRequest("GET", "/vehicles/color/invalid/year/2010", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Grabar la respuesta del handler
	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/vehicles/color/{color}/year/{year}", handler.FindByColorAndYear())
	router.ServeHTTP(rr, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Verificar el cuerpo de la respuesta
	expectedResponse := `{"status":"Internal Server Error","message":"internal error"}`
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Verificar llamada al mock
	mockService.AssertExpectations(t)
}

// Test for SearchByWeightRange: Successful request
func TestHandlerVehicle_SearchByWeightRange(t *testing.T) {
	// Mock del servicio
	mockService := service.NewMocksVehicle()
	handler := NewHandlerVehicle(mockService)

	// Configurar expectativas del mock
	mockService.On("SearchByWeightRange", internal.SearchQuery{
		FromWeight: 1000,
		ToWeight:   2000,
	}, true).Return(map[int]internal.Vehicle{
		1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "Mustang",
				Model:           "model",
				Registration:    "registration",
				Color:           "red",
				FabricationYear: 2010,
				Capacity:        5,
				MaxSpeed:        100,
				FuelType:        "fuelType",
				Transmission:    "transmission",
				Weight:          1500,
				Dimensions: internal.Dimensions{
					Height: 1,
					Length: 1,
					Width:  1,
				},
			},
		},
	}, nil)

	// Crear solicitud HTTP simulada
	req, err := http.NewRequest("GET", "/vehicles/weight?weight_min=1000&weight_max=2000", nil)
	//println("AAAAAAAAAAAAAAA", internal.SearchQuery)
	if err != nil {
		t.Fatal(err)
	}

	// Grabar la respuesta del handler
	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/vehicles/weight", handler.SearchByWeightRange())
	router.ServeHTTP(rr, req)

	// Verificar el código de estado
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verificar el cuerpo de la respuesta
	expectedResponse := `{"message":"vehicles found","data":{}}`
	assert.Equal(t, expectedResponse, rr.Body.String())

	// Verificar llamada al mock
	mockService.AssertExpectations(t)
}
