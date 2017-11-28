package unit

import (
	"reflect"
	"testing"

	"github.com/go-openapi/runtime/middleware"
	"github.com/uhuchain/uhuchain-api/ledger"
	"github.com/uhuchain/uhuchain-api/models"
	"github.com/uhuchain/uhuchain-api/restapi/handler"
	"github.com/uhuchain/uhuchain-api/restapi/operations/car"
)

func TestRequestHandler_HandleAddCar(t *testing.T) {
	payload := `{
		"brand": "Volkswagen",
		"id": 12345,
		"model": "Sharan GTI",
		"policies": [
			{
				"claims": [
					{
						"date": "2016-11-01",
						"description": "Something bad happend",
						"id": 12345
					}
				],
				"endDate": "2017-09-01",
				"id": 12345,
				"insuranceId": 12345,
				"insuranceName": "Zurich Insurance Group",
				"startDate": "2016-09-01"
			}
		],
		"vehicleId": "THK34SDM6A2D34"
	}`

	carPayload := models.Car{}
	carPayload.UnmarshalBinary([]byte(payload))

	mockClient := &ledger.ClientMock{
		QueryResponse: payload,
	}
	requestHandler := handler.NewRequestHandler(mockClient)

	newCarParams := car.NewAddCarParams()
	newCarParams.Body = &carPayload

	type args struct {
		params car.AddCarParams
	}
	tests := []struct {
		name string
		args args
		want middleware.Responder
	}{
		{
			name: "Save car successfully",
			args: args{
				newCarParams,
			},
			want: car.NewAddCarCreated(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestHandler.HandleAddCar(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestHandler.HandleAddCar() = %v, want %v", got, tt.want)
			}
		})
	}
}