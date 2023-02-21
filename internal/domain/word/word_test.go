package word

import (
	"reflect"
	"testing"
)

func TestGenerateFilterWordCloud(t *testing.T) {
	type args struct {
		wordDict FilterCloud
	}
	tests := []struct {
		name     string
		args     args
		want     map[string]int
		original map[string]int
	}{
		{
			name:     "Delete one key",
			args:     args{wordDict: FilterCloud{Filter: []string{"TestKey1"}, WordCloud: map[string]int{"TestKey1": 1, "TestKey2": 2}}},
			want:     map[string]int{"TestKey2": 2},
			original: map[string]int{"TestKey1": 1, "TestKey2": 2},
		},
		{
			name:     "Delete two key",
			args:     args{wordDict: FilterCloud{Filter: []string{"TestKey1", "TestKey2"}, WordCloud: map[string]int{"TestKey1": 1, "TestKey2": 2}}},
			want:     map[string]int{},
			original: map[string]int{"TestKey1": 1, "TestKey2": 2},
		},
		{
			name:     "Filter is null",
			args:     args{wordDict: FilterCloud{Filter: []string{}, WordCloud: map[string]int{"TestKey1": 1, "TestKey2": 2}}},
			want:     map[string]int{"TestKey1": 1, "TestKey2": 2},
			original: map[string]int{"TestKey1": 1, "TestKey2": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateFilterWordCloud(tt.args.wordDict); !reflect.DeepEqual(got, tt.want) || !reflect.DeepEqual(tt.original, tt.args.wordDict.WordCloud) {
				t.Errorf("GenerateFilterWordCloud() = %v, want %v", got, tt.want)
			}
		})
	}
}
