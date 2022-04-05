package currencylayer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	apiURL = `http://apilayer.net/api/`
)

type options struct {
	accessKey  string
	currencies []string
	source     string
	format     string
}

type response struct {
	Success   bool               `json:"success,omitempty"`
	Terms     string             `json:"terms,omitempty"`
	Privacy   string             `json:"privacy,omitempty"`
	Timestamp int                `json:"timestamp,omitempty"`
	Source    string             `json:"source,omitempty"`
	Quotes    map[string]float64 `json:"quotes,omitempty"`
	Error     struct {
		Code int    `json:"code,omitempty"`
		Type string `json:"type,omitempty"`
		Info string `json:"info,omitempty"`
	} `json:"error"`
}

type quotes struct {
	quoteList map[string]float64
}

func New() *options {
	return &options{format: `1`} // by default 1
}

//AddAccessKey adds access key to option struct
func (o *options) AddAccessKey(accessKey string) {
	o.accessKey = accessKey
	o.format = `1`
}

//AddCurrencies add list of currencies to option struct
func (o *options) AddCurrencies(currencies ...string) {
	o.currencies = currencies
}

//AddSource add source currency to option struct, by default is USD
func (o *options) AddSource(source string) {
	o.source = source
}

//EditFormat edit format to option struct
func (o *options) EditFormat(format string) {
	o.format = format
}

//GetQuotes does API query to currencyLayer, default API url http://apilayer.net/api/
func (o *options) GetQuotes(url string) (*quotes, error) {
	if o.accessKey == `` {
		return nil, errors.New(`access key is empty`)
	}

	if url == `` {
		url = apiURL
	}

	resp, err := http.Get(url + `live?` + o.toQueryParams())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(`incorrect response status code`)
	}

	res := new(response)
	if err = json.Unmarshal(raw, res); err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, errors.New(res.Error.Type)
	}

	return &quotes{quoteList: res.Quotes}, nil
}

//QuotesList return quotes list
func (q *quotes) QuotesList() map[string]float64 {
	return q.quoteList
}

//GetQuote return quote by currency pair. Example USDKES
func (q *quotes) GetQuote(quoteName string) (float64, bool) {
	value, found := q.quoteList[strings.ToUpper(quoteName)]
	return value, found
}

//toQueryParams get query representation of option struct
func (o *options) toQueryParams() string {
	return fmt.Sprintf(
		"access_key=%s&currencies=%s&source=%s&format=%s",
		o.accessKey,
		strings.Join(o.currencies, `,`),
		o.source,
		o.format,
	)
}
