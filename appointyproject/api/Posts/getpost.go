package Posts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/mongo/readpref"
)

func doGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	je := json.NewEncoder(w)
	je.Encode(db)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  db = client.Database("instagram").collections("posts")
  collection = db.Collection("posts")
	var posts[] Post
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur,err = collection.Find(ctx,jsonstring{})
	defer cur.close(ctx)
	for cur.Next(ctx){
		var post Post
		cursor.Decode(&post)
		users = append(posts,post)
	}
  // adding pagination feature
	page,_ := strconv.Atoi(ctx.Query("page","1"))
	var perpage int64 = 9
	findOptions.SetSkip((int64(page)-1)*perpage)
	findOptions.setLimit(perpage)

	return JSON(posts)

}
