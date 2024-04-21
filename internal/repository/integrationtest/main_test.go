package integrationtest

import (
	"fmt"
	"kings-comp/internal/repository/redis"
	"kings-comp/pkg/testhelper"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
)

var redisPort string

func TestMain(m *testing.M) {
	fmt.Println("??")
	if !testhelper.IsIntegration() {
		fmt.Println("what the fuck?")
		return
	}

	pool := testhelper.StartDockerPool()

	// set up the redis container for tests
	redisRes := testhelper.StartDockerInstance(pool, "redis/redis-stack-server", "latest",
		func(res *dockertest.Resource) error {
			port := res.GetPort("6379/tcp")
			_, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", port))
			return err
		})

	redisPort = redisRes.GetPort("6379/tcp")

	// now run tests
	exitCode := m.Run()
	redisRes.Close()
	os.Exit(exitCode)
}
