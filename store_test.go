package keyval

import (
	"reflect"
	"testing"
)

func TestMemoryStore_Put(t *testing.T) {
	type args struct {
		key string
		val []byte
	}
	tests := []struct {
		name    string
		store   *MemoryStore
		args    args
		wantErr bool
	}{
		{
			name:  "put a not exist key",
			store: NewMemoryStore(),
			args: args{
				key: "Hello",
				val: []byte("Gopher"),
			},
			wantErr: false,
		},
		{
			name: "put a exist key",
			store: &MemoryStore{
				data: map[string][]byte{"Hello": []byte("World")},
			},
			args: args{
				key: "Hello",
				val: []byte("Gopher"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.store.Put(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("MemoryStore.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryStore_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		store   *MemoryStore
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:  "get a not exist key",
			store: NewMemoryStore(),
			args: args{
				key: "Hello",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "get a exist key",
			store: &MemoryStore{
				data: map[string][]byte{"Hello": []byte("World")},
			},
			args: args{
				key: "Hello",
			},
			want:    []byte("World"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.store.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemoryStore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoryStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryStore_Drop(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		store   *MemoryStore
		args    args
		wantErr bool
	}{
		{
			name:  "drop a not exist key",
			store: NewMemoryStore(),
			args: args{
				key: "Hello",
			},
			wantErr: true,
		},
		{
			name: "drop a exist key",
			store: &MemoryStore{
				data: map[string][]byte{"Hello": []byte("World")},
			},
			args: args{
				key: "Hello",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.store.Drop(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("MemoryStore.Drop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryStore_Keys(t *testing.T) {
	tests := []struct {
		name  string
		store *MemoryStore
		want  []string
	}{
		{
			name:  "empty store",
			store: NewMemoryStore(),
			want:  nil,
		},
		{
			name: "exist key",
			store: &MemoryStore{
				data: map[string][]byte{
					"Hello": []byte("World"),
					"bye":   []byte("(0_0)"),
				},
			},
			want: []string{"Hello", "bye"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.store.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoryStore.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}
