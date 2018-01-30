package models

import (
	"gopkg.in/mgo.v2"
)

var HubDBIndexes = map[string][]mgo.Index{
	"user": []mgo.Index{
		mgo.Index{
			Key:        []string{"nickname"},
			Unique:     true,
			Sparse:     true,
			Background: true,
		},
		mgo.Index{
			Key:        []string{"createtime"},
			Unique:     false,
			Sparse:     false,
			Background: true,
		},
	},
}

func CreateDBIndexes(clusterName string, dbName string, collName string, pos int) (err error) {
	var collNames []string
	if collName == "" {
		collNames, err = DBCollNames(clusterName, dbName)
		if err != nil {
			return err
		}
	} else {
		collNames = []string{collName}
	}

	for _, collName := range collNames {
		coll, err := NewMongoColl(clusterName, dbName, collName)
		if err != nil {
			return err
		}

		for i, index := range HubDBIndexes[collName] {
			if pos == -1 || i == pos {
				err := coll.EnsureIndex(index)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
