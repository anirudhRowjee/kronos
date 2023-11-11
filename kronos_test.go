package main

import "testing"

func TestLogicalClock_Update(t *testing.T) {
	type fields struct {
		Value      int
		Prev_value int
	}
	type args struct {
		new *LogicalClock
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{

		// Check if the basic implementation works
		{"Basic test -> Successful Update", fields{0, 0}, args{&LogicalClock{1, 0}}, 2, false},

		// Ensure that for a lower timestamp recieved, the event is recieved, but the timestamp is not updated
		{"Basic test -> Unsuccessful Update", fields{2, 0}, args{&LogicalClock{1, 0}}, 3, false},

		// TODO add more tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &LogicalClock{
				Value:      tt.fields.Value,
				Prev_value: tt.fields.Prev_value,
			}
			got, err := self.Update(tt.args.new)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogicalClock.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LogicalClock.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
