// parameters, eg: curl 'localhost:8000/?expr=x*y&env=x=3%20y=4'
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/srinav141/go-practice/src/ch7_interfaces/eval"
)

func parsEnv(s string) (eval.Env, error) {
	env := eval.Env{}
	assignments := strings.Fields(s)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			return env, fmt.Errorf("bad assignment %s", a)
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return env, fmt.Errorf("bad value for %s:%s", ident, err)
		}
		env[eval.Var(ident)] = val
	}
	return env, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		exprStr := r.FormValue("expr")
		fmt.Println("Expr is %s", exprStr)
		if exprStr == "" {
			http.Error(w, "no expression", http.StatusBadRequest)
			return
		}
		env, err := parsEnv(r.FormValue("env"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		expr, err := eval.Parse(exprStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, expr.Eval(env))
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
