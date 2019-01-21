package swap

import "testing"

func TestSwap(t *testing.T) {
	type args struct {
		s []interface{}
		i int
		j int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{[]interface{}{1, 2}, 0, 1}, false},
		{"2", args{[]interface{}{1, 'a', "aa"}, 0, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Swap(tt.args.s, tt.args.i, tt.args.j); (err != nil) != tt.wantErr {
				t.Errorf("Swap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsSameSlice(t *testing.T) {
	type args struct {
		a []interface{}
		b []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]interface{}{1},[]interface{}{2}}, false},
		{"2", args{[]interface{}{1},[]interface{}{'a'}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameSlice(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsSameSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
