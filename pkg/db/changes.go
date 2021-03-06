package db

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetChangesList(namespace, resource, collectionName string, limit int, startTime time.Time) ([]model.Event, error) {
	mongo.logger.WithField("collection", collectionName).Debugf("getting changes")
	var collection = mongo.db.C(collectionName)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"resourcename":      resource,
		"dateadded": bson.M{
			"$gte": startTime,
		},
	}).Sort("-eventtime").Limit(limit).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get changes")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetChangesInNamespacesList(collectionName string, limit int, startTime time.Time, namespaces ...string) ([]model.Event, error) {
	mongo.logger.WithField("collection", collectionName).Debugf("getting changes in namespace")
	var collection = mongo.db.C(collectionName)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": bson.M{
			"$in": namespaces,
		},
		"dateadded": bson.M{
			"$gte": startTime,
		},
	}).Sort("-eventtime").Limit(limit).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get changes in namespace")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetAllChangesList(collectionName string, limit int, startTime time.Time) ([]model.Event, error) {
	mongo.logger.WithField("collection", collectionName).Debugf("getting changes in all namespaces")
	var collection = mongo.db.C(collectionName)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"dateadded": bson.M{
			"$gte": startTime,
		},
	}).Sort("-eventtime").Limit(limit).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get changes in all namespaces")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}
