// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user

import (
	"net/http"

	"github.com/harness/gitness/app/api/controller/user"
	"github.com/harness/gitness/app/api/render"
	"github.com/harness/gitness/app/api/request"
	"github.com/harness/gitness/types/enum"
)

// HandleDeleteToken returns an http.HandlerFunc that
// deletes a token of a user.
func HandleDeleteToken(userCtrl *user.Controller, tokenType enum.TokenType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		session, _ := request.AuthSessionFrom(ctx)
		userUID := session.Principal.UID

		tokenUID, err := request.GetTokenUIDFromPath(r)
		if err != nil {
			render.BadRequest(w)
			return
		}

		err = userCtrl.DeleteToken(ctx, session, userUID, tokenType, tokenUID)
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}

		render.DeleteSuccessful(w)
	}
}
