package normalizer_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/danielbintar/go-record/normalizer"
)

type transformSafeName struct {
  origin string
  target string
}

func TestSafeName(t *testing.T) {
  dict := []*transformSafeName{
    &transformSafeName{origin: "", target: ""},
    &transformSafeName{origin: "Product", target: "product"},
    &transformSafeName{origin: "ProductDetail", target: "product_detail"},
    &transformSafeName{origin: "Product20Day", target: "product_20_day"},
  }

  for _, pair := range dict {
    assert.Equal(t, normalizer.SafeName(pair.origin), pair.target)
  }
}
