package graph

import (
	"context"
	"fmt"
	"meetmeup/models"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
)

// create dataloader middleware
// it will be on top of our graphql layer
// it will take userlader from gen
// it will batch fetch and cashes user

const userloaderKey = "userloader"

func DataloaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			// I am giving  a function which will be used to load the users data
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User

				// I am getting ids in string convert it into int
				// in keyword for array of id's and it will find it.
				// intIds := make([]int, 0)
				// for _, i := range ids {
				// 	inte, err := strconv.Atoi(i)

				// 	if err != nil {
				// 		log.Fatal("error", err)
				// 	}

				// 	intIds = append(intIds, inte)
				// }

				err := db.Model(&users).Where("id IN (?)", pg.In(ids)).Select()

				if err != nil {
					fmt.Println("Error in loader", err)
					return nil, []error{err}
				}

				// send users in sync
				fmt.Println("Sync ids", ids)

				u := make(map[string]*models.User, len(users))

				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*models.User, len(ids))

				// ig error here
				for i, id := range ids {
					result[i] = u[id]
				}

				return result, []error{err}
			},
		}

		// add the current context with vale inside request
		ctx := context.WithValue(r.Context(), userloaderKey, &userloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userloaderKey).(*UserLoader)
}
