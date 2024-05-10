package service

import (
	"app/internal"
	"app/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
func TestServiceVehicleDefault_FindAll(t *testing.T) {
	type arrange struct {
		rpMock func() *repository.MockVehicle
	}
	type output struct {
		err     error
		vehicle []internal.Vehicle
	}
	type testCase struct {
		name    string
		arrange arrange
		output  output
	}
	testCases := []testCase{
		{
			name: "success get all vehicles",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {
					rp := repository.NewMocksVehicle()
					vehicles := []internal.Vehicle{
						{
							Id: 1,
							VehicleAttributes: internal.VehicleAttributes{
								Brand:           "brand",
								Model:           "model",
								Registration:    "registration",
								Color:           "color",
								FabricationYear: 2020,
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
						{
							Id: 2,
							VehicleAttributes: internal.VehicleAttributes{
								Brand:           "brand2",
								Model:           "model2",
								Registration:    "registration2",
								Color:           "color2",
								FabricationYear: 2021,
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
					rp.On("FindAll").Return(map[int]internal.Vehicle{}, nil)
					return rp
				},
			},
			output: output{
				err: nil,
				vehicle: []internal.Vehicle{
					{
						Id: 1,
						VehicleAttributes: internal.VehicleAttributes{
							Brand:           "brand",
							Model:           "model",
							Registration:    "registration",
							Color:           "color",
							FabricationYear: 2020,
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
					{
						Id: 2,
						VehicleAttributes: internal.VehicleAttributes{
							Brand:           "brand2",
							Model:           "model2",
							Registration:    "registration2",
							Color:           "color2",
							FabricationYear: 2021,
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
				},
			},
		},
		{
			name: "error get all vehicles",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {
					rp := repository.NewMocksVehicle()
					rp.On("FindAll").Return([]internal.Vehicle{}, internal.ErrRepositoryInvalidFind)
					return rp
				},
			},
			output: output{
				err: internal.ErrRepositoryInvalidFind,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			rpMock := tc.arrange.rpMock()
			sv := NewServiceVehicleDefault(rpMock)

			// act
			vehicles, err := sv.FindAll()

			// assert
			require.Equal(t, tc.output.err, err)
		})

	}
}
*/

func TestServiceVehicleDefault_FindByColorAndYear(t *testing.T) {
	type arrange struct {
		rpMock func() *repository.MockVehicle
	}
	type input struct {
		color string
		year  int
	}
	type output struct {
		err error
	}
	type testCase struct {
		name    string
		arrange arrange
		input   input
		output  output
	}
	testCases := []testCase{
		{
			name: "success get vehicles by color and year",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {
					rp := repository.NewMocksVehicle()
					rp.On("FindByColorAndYear", "red", 2020).Return(map[int]internal.Vehicle{}, nil)
					return rp
				},
			},
			input: input{
				color: "red",
				year:  2020,
			},
			output: output{
				err: nil,
			},
		},
		{
			name: "error get vehicles by color and year",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {
					rp := repository.NewMocksVehicle()
					rp.On("FindByColorAndYear", "red", 2020).Return(map[int]internal.Vehicle{}, internal.ErrRepositoryInvalidFind)
					return rp
				},
			},
			input: input{
				color: "red",
				year:  2020,
			},
			output: output{
				err: internal.ErrRepositoryInvalidFind,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			rpMock := tc.arrange.rpMock()
			sv := NewServiceVehicleDefault(rpMock)

			// act
			_, err := sv.FindByColorAndYear(tc.input.color, tc.input.year)

			// assert
			require.Equal(t, tc.output.err, err)
		})
	}
}

func TestServiceVehicleDefault_FindByBrandAndYearRange(t *testing.T) {
	type arrange struct {
		rpMock func() *repository.MockVehicle
	}
	type input struct {
		brand     string
		startYear int
		endYear   int
	}
	type output struct {
		err     error
		vehicle []internal.Vehicle
	}
	type testCase struct {
		name    string
		arrange arrange
		input   input
		output  output
	}
	testCases := []testCase{
		// Test case 1: Success Query - Find vehicles by brand and year range without errors
		{
			name: "success get vehicles by brand and year range without errors",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {

					vehicles := map[int]internal.Vehicle{
						1: {
							Id: 1,
							VehicleAttributes: internal.VehicleAttributes{
								Brand:           "Mustang",
								Model:           "model",
								Registration:    "registration",
								Color:           "color",
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
								Color:           "color2",
								FabricationYear: 2011,
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
					rp := repository.NewMocksVehicle()
					rp.On("FindByBrandAndYearRange", "Mustang", 2009, 2012).Return(vehicles, nil)
					return rp
				},
			},
			input: input{
				brand:     "Mustang",
				startYear: 2009,
				endYear:   2012,
			},
			output: output{
				err: nil,
				vehicle: []internal.Vehicle{
					{
						Id: 1,
						VehicleAttributes: internal.VehicleAttributes{
							Brand:           "Mustang",
							Model:           "model",
							Registration:    "registration",
							Color:           "color",
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
					{
						Id: 2,
						VehicleAttributes: internal.VehicleAttributes{
							Brand:           "Mustang",
							Model:           "model2",
							Registration:    "registration2",
							Color:           "color2",
							FabricationYear: 2011,
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
				},
			},
		},

		{
			name: "error get vehicles by brand and year range",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {
					rp := repository.NewMocksVehicle()
					rp.On("FindByBrandAndYearRange", "Ferrari", 2050, 2060).Return(map[int]internal.Vehicle{}, internal.ErrRepositoryInvalidFind)
					return rp
				},
			},
			input: input{
				brand:     "Ferrari",
				startYear: 2050,
				endYear:   2060,
			},
			output: output{
				err: internal.ErrRepositoryInvalidFind,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			rpMock := tc.arrange.rpMock()
			sv := NewServiceVehicleDefault(rpMock)

			// act
			result, err := sv.FindByBrandAndYearRange(tc.input.brand, tc.input.startYear, tc.input.endYear)

			// assert
			//require.NoError(t, err) // Verificar que no hay errores

			// Verificar que la longitud de la lista de vehículos sea igual a la longitud de la lista esperada
			require.Len(t, result, len(tc.output.vehicle))

			// Verificar que cada vehículo en el resultado está en la lista de vehículos esperada
			for _, v := range result {
				found := false
				for _, expected := range tc.output.vehicle {
					if v.Id == expected.Id {
						found = true
						break
					}
				}
				require.True(t, found, "Vehicle not found in expected output")
			}

			require.ErrorIs(t, err, tc.output.err)
		})
		/*
			t.Run(tc.name, func(t *testing.T) {
				// arrange
				rpMock := tc.arrange.rpMock()
				sv := NewServiceVehicleDefault(rpMock)

				// act
				result, err := sv.FindByBrandAndYearRange(tc.input.brand, tc.input.startYear, tc.input.endYear)

				// assert
				require.Equal(t, len(tc.output.vehicle), len(result))
				require.Equal(t, tc.output.err, err)
			})
		*/
	}
}

func TestServiceVehicleDefault_AverageCapacityByBrand(t *testing.T) {
	type arrange struct {
		rpMock func() *repository.MockVehicle
	}
	type input struct {
		brand string
	}
	type output struct {
		err error
		avg int
	}
	type testCase struct {
		name    string
		arrange arrange
		input   input
		output  output
	}
	testCases := []testCase{
		// Success finding vehicles by brand
		{
			name: "success get average capacity by brand",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {
					rp := repository.NewMocksVehicle()
					rp.On("FindByBrand", "Ferrari").Return(map[int]internal.Vehicle{
						1: {
							Id: 1,
							VehicleAttributes: internal.VehicleAttributes{
								Brand:           "Ferrari",
								Model:           "model",
								Registration:    "registration",
								Color:           "color",
								FabricationYear: 2020,
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
								Brand:           "Ferrari",
								Model:           "model2",
								Registration:    "registration2",
								Color:           "color2",
								FabricationYear: 2021,
								Capacity:        7,
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
					}, nil)
					return rp
				},
			},
			input: input{
				brand: "Ferrari",
			},
			output: output{
				err: nil,
				avg: 6,
			},
		},

		// Error finding vehicles by brand
		{
			name: "error get average capacity by brand",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {
					rp := repository.NewMocksVehicle()
					rp.On("FindByBrand", "Ferrari").Return(map[int]internal.Vehicle{}, internal.ErrServiceNoVehicles)
					return rp
				},
			},
			input: input{
				brand: "Ferrari",
			},
			output: output{
				err: internal.ErrServiceNoVehicles,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			rpMock := tc.arrange.rpMock()
			sv := NewServiceVehicleDefault(rpMock)

			// act
			avg, err := sv.AverageCapacityByBrand(tc.input.brand)

			// assert
			require.Equal(t, tc.output.avg, avg)
			require.Equal(t, tc.output.err, err)
		})
	}
}

func TestServiceVehicleDefault_AverageMaxSpeedByBrand(t *testing.T) {
	type arrange struct {
		rpMock func() *repository.MockVehicle
	}
	type input struct {
		brand string
	}
	type output struct {
		err error
		avg float64
	}
	type testCase struct {
		name    string
		arrange arrange
		input   input
		output  output
	}
	testCases := []testCase{
		// Success finding vehicles by brand
		{
			name: "success get average max speed by brand",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {
					rp := repository.NewMocksVehicle()
					rp.On("FindByBrand", "Ferrari").Return(map[int]internal.Vehicle{
						1: {
							Id: 1,
							VehicleAttributes: internal.VehicleAttributes{
								Brand:           "Ferrari",
								Model:           "model",
								Registration:    "registration",
								Color:           "color",
								FabricationYear: 2020,
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
								Brand:           "Ferrari",
								Model:           "model2",
								Registration:    "registration2",
								Color:           "color2",
								FabricationYear: 2021,
								Capacity:        7,
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
					}, nil)
					return rp
				},
			},
			input: input{
				brand: "Ferrari",
			},
			output: output{
				err: nil,
				avg: 150,
			},
		},

		// Error finding vehicles by brand
		{
			name: "error get average max speed by brand",
			arrange: arrange{
				rpMock: func() *repository.MockVehicle {
					rp := repository.NewMocksVehicle()
					rp.On("FindByBrand", "Ferrari").Return(map[int]internal.Vehicle{}, internal.ErrServiceNoVehicles)
					return rp
				},
			},
			input: input{
				brand: "Ferrari",
			},
			output: output{
				err: internal.ErrServiceNoVehicles,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			rpMock := tc.arrange.rpMock()
			sv := NewServiceVehicleDefault(rpMock)

			// act
			avg, err := sv.AverageMaxSpeedByBrand(tc.input.brand)

			// assert
			require.Equal(t, tc.output.avg, avg)
			require.Equal(t, tc.output.err, err)
		})
	}
}
