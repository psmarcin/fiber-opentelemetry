package fiber_opentelemetry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_configDefault(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "should set default values for no configs",
			args: args{
				config: []Config{},
			},
			want: Config{
				Tracer:                ConfigDefault.Tracer,
				TracerStartAttributes: ConfigDefault.TracerStartAttributes,
				SpanName:              ConfigDefault.SpanName,
				LocalKeyName:          ConfigDefault.LocalKeyName,
			},
		},
		{
			name: "should set default only for TracerStartAttributes",
			args: args{
				config: []Config{
					{
						Tracer:                ConfigDefault.Tracer,
						TracerStartAttributes: nil,
						SpanName:              ConfigDefault.SpanName,
						LocalKeyName:          ConfigDefault.LocalKeyName,
					},
				},
			},
			want: Config{
				Tracer:                ConfigDefault.Tracer,
				TracerStartAttributes: ConfigDefault.TracerStartAttributes,
				SpanName:              ConfigDefault.SpanName,
				LocalKeyName:          ConfigDefault.LocalKeyName,
			},
		},
		{
			name: "should set default only for SpanName",
			args: args{
				config: []Config{
					{
						Tracer:                ConfigDefault.Tracer,
						TracerStartAttributes: ConfigDefault.TracerStartAttributes,
						SpanName:              "",
						LocalKeyName:          ConfigDefault.LocalKeyName,
					},
				},
			},
			want: Config{
				Tracer:                ConfigDefault.Tracer,
				TracerStartAttributes: ConfigDefault.TracerStartAttributes,
				SpanName:              ConfigDefault.SpanName,
				LocalKeyName:          ConfigDefault.LocalKeyName,
			},
		},
		{
			name: "should set default only for LocalKeyName",
			args: args{
				config: []Config{
					{
						Tracer:                ConfigDefault.Tracer,
						TracerStartAttributes: ConfigDefault.TracerStartAttributes,
						SpanName:              ConfigDefault.SpanName,
						LocalKeyName:          "",
					},
				},
			},
			want: Config{
				Tracer:                ConfigDefault.Tracer,
				TracerStartAttributes: ConfigDefault.TracerStartAttributes,
				SpanName:              ConfigDefault.SpanName,
				LocalKeyName:          ConfigDefault.LocalKeyName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := configDefault(tt.args.config...)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
