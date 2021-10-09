package users

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
  aUser:= &User{}
  err:=jd.Decode(aUser)
  if nil != err {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  //adding locking feature to prevent logging in at the same time
  lock.lock()
  nextUserId++
  aUser.ID = nextUserId
  db = append(db,aUser)
  lock.Unlock()
  respUser: User{ID: 1,Username:"Sreejan",Email:"abcd@gmail.com",Password:"abcdf"}
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  db = client.Database("instagram").collections("users")
  collection = db.Collection("users")
  res, err := collection.InsertOne(ctx, respUser)
  id := res.InsertedID
  je := json.NewEncoder(w)
  je.Encode(respUser)
}
