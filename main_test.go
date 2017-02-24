package json

import "testing"

func Test(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{`s:5:"aaaaa"`, `"aaaaa"`},
		{"", `""`},
		{`a:2:{i:0;s:70:"http://st-im.kinopoisk.ru/im/kadr/5/3/9/kinopoisk.ru-Kitaro-539766.jpg";i:1;s:70:"http://st-im.kinopoisk.ru/im/kadr/5/3/9/kinopoisk.ru-Kitaro-539765.jpg";}`, `{"0":"http://st-im.kinopoisk.ru/im/kadr/5/3/9/kinopoisk.ru-Kitaro-539766.jpg","1":"http://st-im.kinopoisk.ru/im/kadr/5/3/9/kinopoisk.ru-Kitaro-539765.jpg"}`},
		{`s:5:"яяяяя"`, `"яяяяя"`},
		{`s:6:"яzяяяя"`, `"яzяяяя"`},
		{`i:0;`, `"0"`},
	}
	for _, c := range cases {
		got, err := DecodeToJSON(c.in)
		if err != nil {
			t.Errorf("Error:", err)
		}
		if got != c.want {
			t.Errorf("Decode(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}
