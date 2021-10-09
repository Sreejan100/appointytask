package Posts

import(
  "encoding/json"
  "net/http"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/mongo/readpref"
)

func doPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json")
  jd := json.NewDecoder(r.GetBody)
  aPost:= &Post{}
  err:=jd.Decode(aPost)
  if nil != err {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  // code for preventing deadlock when multiple posts are being posted at the same time
  lock.lock()
  nextPostId++
  aPost.ID = nextPostId
  db = append(db,aPost)
  lock.Unlock()
  respPosts: Post{ID: 1,Caption:"A place to visit",Image:"http://google.com/Hyderabad",Time:"19:59PM 8 October 2021"}
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  db = client.Database("instagram").collections("posts")
  collection = db.Collection("posts")
  res, err := collection.InsertOne(ctx, respPosts)
  id := res.InsertedID
  je := json.NewEncoder(w)
  je.Encode(respPosts)
}
