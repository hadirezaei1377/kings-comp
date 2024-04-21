package integrationtest

import (
	"context"
	"fmt"
	"kings-comp/internal/entity"
	"kings-comp/internal/repository"
	"kings-comp/internal/repository/redis"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testType struct {
	ID   string
	Name string
}

func (t testType) EntityID() entity.ID {
	return entity.NewID("testType", t.ID)
}

func TestCommonBehaviourSetAndGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	ctx := context.Background()

	cb := repository.NewRedisCommonBehaviour[testType](redisClient)
	err = cb.Save(ctx, &testType{
		ID:   "12",
		Name: "Sajad Jalilian",
	})
	assert.NoError(t, err)

	err = cb.Save(ctx, &testType{
		ID:   "13",
		Name: "Amirreza",
	})
	assert.NoError(t, err)

	val, err := cb.Get(ctx, entity.NewID("testType", "12"))
	assert.NoError(t, err)
	assert.Equal(t, "Sajad Jalilian", val.Name)
	assert.Equal(t, "12", val.ID)

	val, err = cb.Get(ctx, entity.NewID("testType", "13"))
	assert.NoError(t, err)
	assert.Equal(t, "Amirreza", val.Name)
	assert.Equal(t, "13", val.ID)

	err = cb.Save(ctx, &testType{
		ID:   "13",
		Name: "yasin",
	})

	assert.NoError(t, err)
	val, err = cb.Get(ctx, entity.NewID("testType", "13"))
	assert.NoError(t, err)
	assert.Equal(t, "yasin", val.Name)

	val, err = cb.Get(ctx, entity.NewID("testType", "14"))
	assert.ErrorIs(t, repository.ErrNotFound, err)
}
