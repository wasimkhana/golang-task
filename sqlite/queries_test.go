package sqlite

import (
	"reflect"
	"testing"
	"webserver/model"
)

func Test_client_GetHaisByID(t *testing.T) {

	hais := &model.Hais{
		ID: 3,
	}

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Hais
		wantErr bool
	}{
		{
			name:    "GetHaisByID - Successfully Done!",
			args:    args{id: hais.ID},
			want:    hais,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHaisByID(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetHaisByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got[0].ID, tt.want.ID) {
				t.Errorf("GetHaisByID() got = %v, want %v", 3, tt.want)
			}
		})
	}
}

func TestGetHaiCounter(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "GetHaiCounter - Successfully Done!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetHaiCounter()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHaiCounter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetHaiByCityAndHai(t *testing.T) {
	hais := &model.Hais{
		City:    "الرياض",
		HaiName: "ابو طليحة",
	}

	type args struct {
		city string
		hai  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "GetHaiByCityAndHai - Successfully Done!",
			args:    args{city: hais.City, hai: hais.HaiName},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetHaiByCityAndHai(tt.args.city, tt.args.hai)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHaiByCityAndHai() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetHaiByHai(t *testing.T) {
	hais := &model.Hais{
		HaiName: "ابو طليحة",
	}
	type args struct {
		hai string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetHaiByHai - Successfully Done!",
			args: args{hai: hais.HaiName},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetHaiByHai(tt.args.hai)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHaiByHai() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestGetHaiByCity(t *testing.T) {
	hais := &model.Hais{
		City: "الرياض",
	}
	type args struct {
		city string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "GetHaiByCity - Successfully Done!",
			args:    args{city: hais.City},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetHaiByCity(tt.args.city)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHaiByCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
