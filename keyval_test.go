package keyval

import (
	"reflect"
	"sort"
	"testing"
)

func TestKeyval_Put(t *testing.T) {
	type args struct {
		key string
		val []byte
	}
	tests := []struct {
		name    string
		store   Store
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
			name:  "put with the empty key",
			store: NewMemoryStore(),
			args: args{
				key: "",
				val: []byte("Gopher"),
			},
			wantErr: true,
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
			kv := New(tt.store)
			if err := kv.Put(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Keyval.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeyval_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		store   Store
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
			kv := New(tt.store)
			got, err := kv.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keyval.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keyval.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyval_GetString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		store   Store
		args    args
		want    string
		wantErr bool
	}{
		{
			name:  "get a not exist key",
			store: NewMemoryStore(),
			args: args{
				key: "Hello",
			},
			want:    "",
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
			want:    "World",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := New(tt.store)
			got, err := kv.GetString(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keyval.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Keyval.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestKeyval_Drop(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		store   Store
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
			kv := New(tt.store)
			if err := kv.Drop(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Keyval.Drop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeyval_DropAll(t *testing.T) {
	tests := []struct {
		name  string
		store Store
	}{
		{
			name:  "empty",
			store: NewMemoryStore(),
		},
		{
			name: "a exist key",
			store: &MemoryStore{
				data: map[string][]byte{"Hello": []byte("World")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := New(tt.store)
			kv.DropAll()
			if len(kv.Keys()) != 0 {
				t.Errorf("Keyval.DropAll() does not drop")
			}
		})
	}
}

func TestKeyval_Keys(t *testing.T) {
	tests := []struct {
		name  string
		store Store
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
			kv := New(tt.store)
			got := kv.Keys()
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keyval.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyval_PutString(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name    string
		store   Store
		args    args
		wantErr bool
	}{
		{
			name:  "put a not exist key",
			store: NewMemoryStore(),
			args: args{
				key: "Hello",
				val: "Gopher",
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
				val: "Gopher",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := New(tt.store)
			if err := kv.PutString(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Keyval.PutString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeyval_Has(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		store Store
		args  args
		want  bool
	}{
		{
			name:  "empty store",
			store: NewMemoryStore(),
			args: args{
				key: "Hello",
			},
			want: false,
		},
		{
			name: "exist key",
			store: &MemoryStore{
				data: map[string][]byte{"Hello": []byte("World")},
			},
			args: args{
				key: "Hello",
			},
			want: true,
		},
		{
			name: "not exist key",
			store: &MemoryStore{
				data: map[string][]byte{"Hello": []byte("World")},
			},
			args: args{
				key: "Bye",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := New(tt.store)
			if got := kv.Has(tt.args.key); got != tt.want {
				t.Errorf("Keyval.Has() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
