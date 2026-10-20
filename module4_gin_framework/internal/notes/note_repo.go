package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// repo -> data access layer

type Repo struct {
	coll *mongo.Collection 
}

func NewRepo(db *mongo.Database) *Repo{
	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context , note Note)  (Note,error){
	opCtx , cancel := context.WithTimeout(ctx,5*time.Second)

	defer cancel()

	_,err := r.coll.InsertOne(opCtx,note)
	if err != nil {
		return Note{},fmt.Errorf("Insert Note Failed")
	}

	return note,nil 
}

func (r *Repo) List (ctx context.Context) ([]Note,error){
	optCtx, cancel := context.WithTimeout(ctx,5*time.Second)
	defer cancel()

	filter := bson.M{}  // match all docs 

	cursor,err := r.coll.Find(optCtx,filter)
	if err != nil {
		return nil , fmt.Errorf("Find Noted Failed")
	}

	defer cursor.Close(optCtx) // closing cursor after use  , to avoid leaks

	var notes []Note 

	if err := cursor.All(optCtx,&notes); err != nil {
		return nil , fmt.Errorf("Decoding Notes Error")
	}
	return notes ,err 
}

func (r *Repo) GetById(ctx context.Context, id bson.ObjectID) (Note, error){
	optCtx , cancel := context.WithTimeout(ctx,5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id,}

	var note Note 

	err := r.coll.FindOne(optCtx,filter,options.FindOne()).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("Error in finding Note by ID")
	}
	return note,nil 
}


func (r *Repo) UpdateById(ctx context.Context, id bson.ObjectID , req UpdateNoteRequest) (Note ,error){
	optCtx , cancel := context.WithTimeout(ctx,5*time.Second)
	defer cancel()

	filter := bson.M{"_id":id,}
	update := bson.M{
		"$set":bson.M{
			"title":req.Title,
			"content":req.Content,
			"pinned":req.Pinned,
			"updatedAt":time.Now().UTC(),
		},
	}

	after := options.After
	opts := options.FindOneAndUpdate().SetReturnDocument(after)

	var updated Note
	err := r.coll.FindOneAndUpdate(optCtx,filter,update,opts).Decode(&updated)
	if err != nil {
		return Note{},fmt.Errorf("Update Note Failed")
	}
	return updated,nil 
}

func (r *Repo) DeleteById(ctx context.Context, id bson.ObjectID ) (bool ,error) {
	optCtx, cancel := context.WithTimeout(ctx,5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id,}

	res , err := r.coll.DeleteOne(optCtx,filter)
	if err != nil {
		return false,fmt.Errorf("Failed to delete id")
	}

	if res.DeletedCount == 0 {
		return false, nil 
	}

	return true, nil
}