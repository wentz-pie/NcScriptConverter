package alterationncscript

import (
	"reflect"
	"testing"
)

func TestNewConvertedNcScript(t *testing.T) {
	tests := []struct {
		name string
		want *ConvertedNcScript
	}{
		{
			name: "正常系_オブジェクト生成できること",
			want: new(ConvertedNcScript),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConvertedNcScript(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConvertedNcScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertedNcScript_Convert(t *testing.T) {
	type args struct {
		source []string
	}
	tests := []struct {
		name    string
		c       *ConvertedNcScript
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "正常系_カッタースクリプトが変換されること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4701",
					"(T16)",
					"(S4500)",
					"X0.Y0.",
					"G90X0.Y0.",
					"G54",
					"X0.Y0.",
					"M99",
					"%",
				},
			},
			want: []string{
				"",
				"(O4701)",
				"T16",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H16",
				"M01",
				"S4500M3",
				"M8",
				"G05.1Q1",
				"X0.Y0.",
				"G90X0.Y0.",
				"G49",
				"G54",
				"X0.Y0.",
				"G05.1Q0",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_T1桁のカッタースクリプトが変換されること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4701",
					"(T6)",
					"(S4500)",
					"X0.Y0.",
					"G90X0.Y0.",
					"G54",
					"X0.Y0.",
					"M99",
					"%",
				},
			},
			want: []string{
				"",
				"(O4701)",
				"T6",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H6",
				"M01",
				"S4500M3",
				"M8",
				"G05.1Q1",
				"X0.Y0.",
				"G90X0.Y0.",
				"G49",
				"G54",
				"X0.Y0.",
				"G05.1Q0",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_センタードリルスクリプトが変換されること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4702",
					"(T12)",
					"(S2000)",
					"(G82)",
					"X0.Y0.",
					"G90",
					"X0.Y0.",
					"M99",
					"%",
				},
			},
			want: []string{
				"",
				"(O4702)",
				"T12",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H12",
				"M01",
				"S2000M3",
				"M8",
				"G98G82R2.0Z-1.0Q2.0P500F180L0",
				"X0.Y0.",
				"G90",
				"X0.Y0.",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_下穴ドリルスクリプトが変換されること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4702",
					"(T13)",
					"(S1800)",
					"(G83)",
					"X0.Y0.",
					"G90",
					"X0.Y0.",
					"M99",
					"%",
				},
			},
			want: []string{
				"",
				"(O4702)",
				"T13",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H13",
				"M01",
				"S1800M3",
				"M8",
				"G98G83R2.0Z-39.0Q2.0F180L0",
				"X0.Y0.",
				"G90",
				"X0.Y0.",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_リーマスクリプトが変換されること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4702",
					"(T15)",
					"(S1500)",
					"(G85)",
					"X0.Y0.",
					"G90",
					"X0.Y0.",
					"M99",
					"%",
				},
			},
			want: []string{
				"M00",
				"",
				"(O4702)",
				"T15",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H15",
				"M01",
				"S1500M3",
				"M8",
				"G98G85R2.0Z-39.0F150L0",
				"X0.Y0.",
				"G90",
				"X0.Y0.",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_M30の変換がされること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4701",
					"(T16)",
					"(S4500)",
					"X0.Y0.",
					"G90X0.Y0.",
					"X0.Y0.",
					"M30",
					"%",
				},
			},
			want: []string{
				"",
				"(O4701)",
				"T16",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H16",
				"M01",
				"S4500M3",
				"M8",
				"G05.1Q1",
				"X0.Y0.",
				"G90X0.Y0.",
				"X0.Y0.",
				"M09",
				"G91G0G28Z0",
				"G91G0G28B0",
				"G91G0G28C0",
				"(M30)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_タップスクリプトが変換されること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4702",
					"(T15)",
					"(S1500)",
					"(G84)",
					"X0.Y0.",
					"G90",
					"X0.Y0.",
					"M99",
					"%",
				},
			},
			want: []string{
				"M00",
				"",
				"(O4702)",
				"T15",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H15",
				"M01",
				"S1500M3",
				"M8",
				"G98G84R5.0Z-35.0F350L0",
				"X0.Y0.",
				"G90",
				"X0.Y0.",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_回転2桁が変換されること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4701",
					"(T16)",
					"(S45)",
					"X0.Y0.",
					"G90X0.Y0.",
					"G54",
					"X0.Y0.",
					"M99",
					"%",
				},
			},
			want: []string{
				"",
				"(O4701)",
				"T16",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H16",
				"M01",
				"S45M3",
				"M8",
				"G05.1Q1",
				"X0.Y0.",
				"G90X0.Y0.",
				"G49",
				"G54",
				"X0.Y0.",
				"G05.1Q0",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_回転3桁が変換されること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4701",
					"(T16)",
					"(S450)",
					"X0.Y0.",
					"G90X0.Y0.",
					"G54",
					"X0.Y0.",
					"M99",
					"%",
				},
			},
			want: []string{
				"",
				"(O4701)",
				"T16",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H16",
				"M01",
				"S450M3",
				"M8",
				"G05.1Q1",
				"X0.Y0.",
				"G90X0.Y0.",
				"G49",
				"G54",
				"X0.Y0.",
				"G05.1Q0",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
		{
			name: "正常系_面取りクリプトが変換されること",
			c:    new(ConvertedNcScript),
			args: args{
				[]string{
					"%",
					"O4702",
					"(T15)",
					"(S1500)",
					"(G81)",
					"X0.Y0.",
					"G90",
					"X0.Y0.",
					"M99",
					"%",
				},
			},
			want: []string{
				"",
				"(O4702)",
				"T15",
				"M6 Q0",
				"G91G0G28Z0",
				"G54",
				"G90G0X0Y0",
				"G0B0C0",
				"G0W0",
				"G43Z100.H15",
				"M01",
				"S1500M3",
				"M8",
				"G98G81R2.0Z-8.5F200L0",
				"X0.Y0.",
				"G90",
				"X0.Y0.",
				"M5",
				"M9",
				"G91G0G28Z0",
				"(M99)",
				"",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Convert(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertedNcScript.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertedNcScript.Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
