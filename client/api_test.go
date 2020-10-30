package client

import (
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type DuckDuckGoServer struct{}

func (d *DuckDuckGoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["q"]
	if query[0] == "antani" {
		w.Header().Set("Content-Type", "text/html")
		data, err := ioutil.ReadFile("test_data/antani.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(data))
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func TestDuckDuckGoSearchClient_SearchLimited(t *testing.T) {
	type args struct {
		query string
		limit int
	}
	tests := []struct {
		name    string
		args    args
		want    []Result
		wantErr bool
	}{
		{
			name: "shouldParseResponse",
			args: args{
				query: "antani",
				limit: 3,
			},
			want: []Result{
				{
					HtmlFormattedUrl: "www.lettera43.it/howto/come-se-fosse-antani-significato/",
					HtmlTitle:        "Come se fosse <b>antani</b> significato - Lettera43 Come Fare",
					HtmlSnippet:      "Come se fosse <b>antani</b> significato Come se fosse <b>antani</b>: la storia della supercazzola e di come sia diventata di uso comuneLa &#34;supercazzola&#34; del film Amici Miei è ormai di uso comune nelle burle e...",
					FormattedUrl:     "www.lettera43.it/howto/come-se-fosse-antani-significato/",
					Title:            "Come se fosse antani significato - Lettera43 Come Fare",
					Snippet:          "Come se fosse antani significato Come se fosse antani: la storia della supercazzola e di come sia diventata di uso comuneLa \"supercazzola\" del film Amici Miei è ormai di uso comune nelle burle e...",
					Icon: Icon{
						Src:    "//external-content.duckduckgo.com/ip3/www.lettera43.it.ico",
						Width:  16,
						Height: 16,
					},
				},
				{
					HtmlFormattedUrl: "it.wikipedia.org/wiki/Supercazzola",
					HtmlTitle:        "Supercazzola - Wikipedia",
					HtmlSnippet:      "Come se fosse <b>antani</b> anche per lei soltanto in due, oppure in quattro anche scribàcchi confaldina? Come antifurto, per esempio. Vigile: Ma che antifurto, mi faccia il piacere! Questi signori qui stavano sonando loro. &#39;Un s&#39;intrometta!",
					FormattedUrl:     "it.wikipedia.org/wiki/Supercazzola",
					Title:            "Supercazzola - Wikipedia",
					Snippet:          "Come se fosse antani anche per lei soltanto in due, oppure in quattro anche scribàcchi confaldina? Come antifurto, per esempio. Vigile: Ma che antifurto, mi faccia il piacere! Questi signori qui stavano sonando loro. 'Un s'intrometta!",
					Icon: Icon{
						Src:    "//external-content.duckduckgo.com/ip3/it.wikipedia.org.ico",
						Width:  16,
						Height: 16,
					},
				},
				{
					HtmlFormattedUrl: "www.ilsicilia.it/come-se-fosse-antani/",
					HtmlTitle:        "Come se fosse <b>Antani</b> | ilSicilia.it :ilSicilia.it",
					HtmlSnippet:      "Cos&#39;è <b>Antani</b>, come se fosse <b>Antani</b>? È una frase senza senso che si fa beffa dell&#39;autorità. Con tono magniloquente, irride, graffia, scortica concetti e situazioni. <b>Antani</b> non ha nessun senso, nessun significato, come la parola Dada del famoso gruppo di artisti d&#39;Avanguardia, chiamati appunto dadaisti.",
					FormattedUrl:     "www.ilsicilia.it/come-se-fosse-antani/",
					Title:            "Come se fosse Antani | ilSicilia.it :ilSicilia.it",
					Snippet:          "Cos'è Antani, come se fosse Antani? È una frase senza senso che si fa beffa dell'autorità. Con tono magniloquente, irride, graffia, scortica concetti e situazioni. Antani non ha nessun senso, nessun significato, come la parola Dada del famoso gruppo di artisti d'Avanguardia, chiamati appunto dadaisti.",
					Icon: Icon{
						Src:    "//external-content.duckduckgo.com/ip3/www.ilsicilia.it.ico",
						Width:  16,
						Height: 16,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "shouldErrorWhenServerError",
			args: args{
				query: "lampredotto",
				limit: 5,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(&DuckDuckGoServer{})
			defer server.Close()

			c := DuckDuckGoSearchClient{
				baseUrl: server.URL,
			}
			got, err := c.SearchLimited(tt.args.query, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchLimited() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
