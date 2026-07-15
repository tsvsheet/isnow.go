package server

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/uplang/isnow.go/internal/domain"
)

// maxN caps derivation requests.
const maxN = 1000

// deriveBudget bounds the CPU a single derivation request may spend, so a
// pathological pattern (e.g. an impossible bounded window) cannot pin a core.
const deriveBudget = 5 * time.Second

func (s *Server) handleIs(w http.ResponseWriter, r *http.Request) {
	p, ok := parsePattern(w, isnowOf(r))
	if !ok {
		return
	}
	at, ok := s.instantOr400(w, r, "at")
	if !ok {
		return
	}
	if p.Holds(at) {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusPreconditionFailed)
}

func (s *Server) handleCheck(w http.ResponseWriter, r *http.Request) {
	src := isnowOf(r)
	p, ok := parsePattern(w, src)
	if !ok {
		return
	}
	at, ok := s.instantOr400(w, r, "at")
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, checkBody{
		Isnow: src, Canonical: p.Canonical(), At: at.Format(time.RFC3339), Holds: p.Holds(at),
	})
}

func (s *Server) handleNext(w http.ResponseWriter, r *http.Request) { s.derive(w, r, true) }
func (s *Server) handlePrev(w http.ResponseWriter, r *http.Request) { s.derive(w, r, false) }

func (s *Server) derive(w http.ResponseWriter, r *http.Request, forward bool) {
	src := isnowOf(r)
	from, ok := s.instantOr400(w, r, "from")
	if !ok {
		return
	}
	n, ok := countOr400(w, r)
	if !ok {
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), deriveBudget)
	defer cancel()
	occ, err := domain.Derive(ctx, src, from, n, forward)
	if err != nil {
		writeDeriveError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, occurrencesBody{
		Isnow: src, Canonical: canonOf(src), Occurrences: formatInstants(occ),
	})
}

// writeDeriveError maps a cancelled/timed-out search to 504 and a parse fault
// to 400.
func writeDeriveError(w http.ResponseWriter, err error) {
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		writeError(w, http.StatusGatewayTimeout, "timeout", "derivation exceeded the time budget")
		return
	}
	writeError(w, http.StatusBadRequest, isnowCode(err), err.Error())
}

func (s *Server) handleCanon(w http.ResponseWriter, r *http.Request) {
	src := isnowOf(r)
	p, ok := parsePattern(w, src)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, describeBody{Isnow: src, Canonical: p.Canonical()})
}

func (s *Server) handleExplain(w http.ResponseWriter, r *http.Request) {
	src := isnowOf(r)
	p, ok := parsePattern(w, src)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, describeBody{
		Isnow: src, Canonical: p.Canonical(), Explanation: p.Explain(),
	})
}

func (s *Server) handleBuild(w http.ResponseWriter, r *http.Request) {
	v, src, err := domain.Build(buildFieldsOf(r))
	if err != nil {
		writeError(w, http.StatusBadRequest, isnowCode(err), err.Error())
		return
	}
	writeJSON(w, http.StatusOK, describeBody{Isnow: src, Canonical: v.Canonical, Explanation: v.Explain})
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "now": s.now().Format(time.RFC3339)})
}

func (s *Server) instantOr400(w http.ResponseWriter, r *http.Request, key string) (time.Time, bool) {
	at, err := s.instant(r, key)
	if err != nil {
		writeError(w, http.StatusBadRequest, "parameter", err.Error())
		return time.Time{}, false
	}
	return at, true
}

func countOr400(w http.ResponseWriter, r *http.Request) (int, bool) {
	raw := r.URL.Query().Get("n")
	if raw == "" {
		return 1, true
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n < 1 || n > maxN {
		writeError(w, http.StatusBadRequest, "parameter", "n must be 1..1000")
		return 0, false
	}
	return n, true
}

func buildFieldsOf(r *http.Request) domain.BuildFields {
	q := r.URL.Query()
	return domain.BuildFields{
		Year: q.Get("year"), Month: q.Get("month"), Day: q.Get("day"),
		Weekday: q.Get("weekday"),
		Hour:    q.Get("hour"), Minute: q.Get("minute"), Second: q.Get("second"),
		Since: q.Get("since"), Until: q.Get("until"),
	}
}

func formatInstants(ts []time.Time) []string {
	out := make([]string, len(ts))
	for i, t := range ts {
		out[i] = t.Format(time.RFC3339)
	}
	return out
}
