package test

import (
	"HiChat/common"
	"fmt"
	"testing"
)

func TestCrateToken(t *testing.T) {

	type Us struct {
		Name string `json:"name"`
	}

	s := Us{Name: "12312"}
	fmt.Println(common.GenerateTaken(s))
}

func TestParseToken(t *testing.T) {
	s := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.e30.jzJNGpY-XSaZVphG93ti50p_9EloefDWnZNBVKicPvnMd2nJyNrpycGUN3yVLIkXO7YmTFhmhyoRUwRTZlmVc-M9Ht7vyNrK2I3h07EGmvtPUurQPpyj1Ksg9q7EPVCoKRytDQHVaaFp2VQnwUKpwUybOGyjsHEykR8CG_OukuuDdHqdPCwHvgWyZe-vpNkuIh5VpLliZo_7Dm3GwMW062F92r1hweNcHX-6QkENuhYRlL8s3uy0eNcD620d-if5O_dhV4JiZdw8wex4iPPivEYANICiJfJ6UdR1k9vDmyf129Wr3bCZdnrmM_XyG2jg9cAcQwMVivpGWgQYEqbgmQ"
	fmt.Println(common.DecryptTaken(s))
}
