package groupRepository

import (
	"context"
	"log"

	"Golang-API/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	projectId      string = "vibeshare-c2a22"
	collectionName string = "groups"
)

type GroupRepository interface {
	Save(group *entity.Group) (*entity.Group, error)
	FindAll() ([]entity.Group, error)
}

type groupRepo struct{}

// newGroupPostRepository
func NewGroupRepository() GroupRepository {
	return &groupRepo{}
}

func (*groupRepo) Save(group *entity.Group) (*entity.Group, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"GroupID": group.GroupID,
		"Users":   group.Users,
	})

	if err != nil {
		log.Fatalf("Failed addding a new group: %v", err)
		return nil, err
	}
	return group, nil
}

func (*groupRepo) FindAll() ([]entity.Group, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	var groups []entity.Group
	itr := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the groups: %v", err)
			return nil, err
		}
		userInterfaces := doc.Data()["Users"].([]interface{})
		var users []string
		for _, u := range userInterfaces {
			users = append(users, u.(string))
		}
		group := entity.Group{
			GroupID: doc.Data()["GroupID"].(string),
			Users:   users,
		}
		groups = append(groups, group)
	}
	return groups, nil
}
