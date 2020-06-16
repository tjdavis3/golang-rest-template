package api

import (
	"io"
	"net/http"
)

// func (s *server) ReadUser(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	log.Ctx(ctx).Debug().Msg("ReadUser")
// 	uid := chi.URLParam(r, "userID")
// 	u, err := s.db.User.Query().Where(user.UserIDEQ(uid)).Only(ctx)
// 	if err != nil {
// 		render.Render(w, r, ErrServerError(r, err))
// 		return
// 	}
// 	render.JSON(w, r, u)
// 	return
// }
func (s *server) GetTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "success\n")
}
