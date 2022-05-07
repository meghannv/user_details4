package mongo

import (
	"vendor.lib/tng/tng-lib/db/mgo"
)

type Mongo struct {
	mgo.Mongo
	Collection string
}

func (ss *Mongo) FindUser(userId string, ctx context.Context) (model.User, error){
	var user model.User
	collection := ss.Database(ctx, bson.M{"userId":userId}).Decode(&user)
	if err != nil{
		return user, nil
	}
	return user, nil
}

func (ss *Mongo) InjectUser(userr model.User, ctx context.Context) (string, error){
	collection := ss.Database.Collection(ss.Collection["userDetails"])
	filter := bson.M{"userId":userDetails.userId}
	update := bson.M{"$set": &userDetails}
	operation := options.Update().SetUpsert(true)

	result, err := collection.UpdateOne(ctx, filter, update, operation)
	if err != nil{
		return "", err
	}
	return "user details inserted" ,nil
}
