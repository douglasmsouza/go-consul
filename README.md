# Go Consul
A simple helper for accessing [Consul API](https://github.com/hashicorp/consul).

# Usage
Retrieving any available instance of a service

```go
import (
	"fmt"
	"go-consul/consul"
)

func main() {
	c, _ := consul.NewConsulClient("localhost", 8500)
	s, _ := c.GetAvailableService("service-name-here", "")
	fmt.Printf("%s/%d", s.Service.Address, s.Service.Port)
}
```

Retrieving any available URL of a service

```go
import (
	"fmt"
	"go-consul/consul"
)

func main() {
	c, _ := consul.NewConsulClient("localhost", 8500)
	u, _ := c.GetAvailableUrl("service-name-here", "")
	fmt.Printf("%s", u)
}
```

Retrieving any available URL of a service or a default if no one exists

```go
import (
	"fmt"
	"go-consul/consul"
)

func main() {
	c, _ := consul.NewConsulClient("localhost", 8500)
	u, _ := c.GetAvailableUrlOrDefault("service-name-here", "", "http://localhost:8080")
	fmt.Printf("%s", u)
}
```
