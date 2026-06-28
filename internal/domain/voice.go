package domain

import "strings"

type Locale struct {
	Lang    string
	Country string
}

func (l Locale) String() string {
	if l.Lang == "" && l.Country == "" {
		return ""
	}
	return l.Lang + "_" + l.Country
}

type Voice struct {
	Name        string
	Locale      Locale
	Description string
}

type VoiceList []Voice

func (vl VoiceList) FilterByQuery(q string) VoiceList {
	q = strings.TrimSpace(q)
	if q == "" {
		return vl
	}
	if strings.Contains(q, "_") {
		return vl.filter(func(v Voice) bool {
			return strings.EqualFold(v.Locale.String(), q)
		})
	}
	byLang := vl.filter(func(v Voice) bool { return strings.EqualFold(v.Locale.Lang, q) })
	if len(byLang) > 0 {
		return byLang
	}
	return vl.filter(func(v Voice) bool { return strings.EqualFold(v.Locale.Country, q) })
}

func (vl VoiceList) Primary() (Voice, bool) {
	for _, v := range vl {
		if !strings.Contains(v.Name, "(") {
			return v, true
		}
	}
	if len(vl) > 0 {
		return vl[0], true
	}
	return Voice{}, false
}

func (vl VoiceList) filter(pred func(Voice) bool) VoiceList {
	var out VoiceList
	for _, v := range vl {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}
