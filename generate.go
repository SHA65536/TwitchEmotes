package twitchemotes

func (es *EmoteScraper) Generate() []string {
	var res = make([]string, 10)
	for i := range res {
		addOne(&es.Current)
		res[i] = string(es.Current)
	}
	return res
}

func addOne(in *[]byte) {
	for i := len(*in) - 1; i >= 0; i-- {
		if (*in)[i] == '9' {
			(*in)[i] = '0'
			continue
		}
		// If it's not the max, increment and stop
		(*in)[i]++
		return
	}
	*in = append([]byte{'1'}, *in...)
}
