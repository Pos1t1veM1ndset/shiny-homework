package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectedErr error
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde", expectedErr: nil},
		{input: "abccd", expected: "abccd", expectedErr: nil},
		{input: "", expected: "", expectedErr: nil},
		{input: "aaa0b", expected: "aab", expectedErr: nil},
		{input: "\n3a3", expected: "\n\n\naaa", expectedErr: nil},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`, expectedErr: nil},
		{input: `qwe\45`, expected: `qwe44444`, expectedErr: nil},
		{input: `qwe\\5`, expected: `qwe\\\\\`, expectedErr: nil},
		{input: `qwe\\\3`, expected: `qwe\3`, expectedErr: nil},
		{input: `тес3тКириллиц2ы2`, expected: `тессстКириллиццыы`, expectedErr: nil},
		{input: `Бу2кваВ2Конце`, expected: `БуукваВВКонце`, expectedErr: nil},
		{input: `1Ошибка`, expected: ``, expectedErr: ErrInvalidString},
		{input: `日本語が分かりますか。`, expected: `日本語が分かりますか。`, expectedErr: nil},
		{input: `日2本が綺麗ですが2`, expected: `日日本が綺麗ですがが`, expectedErr: nil},
		{input: `\で2す`, expected: `でです`, expectedErr: nil},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			//require.NoError(t, err)
			require.Equal(t, tc.expected, result)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
