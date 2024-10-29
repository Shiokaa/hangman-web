package rank

import "net/http"

func RankRouteur() {

	http.HandleFunc("/rank", rank)
}
