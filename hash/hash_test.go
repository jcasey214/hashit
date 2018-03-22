package hash

import "testing"

func TestCreate(t *testing.T) {
    t.Run("returns a base64 encoded string of the sha512 password", func(t *testing.T){
        actual := Create("angryMonkey")
        expected := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="

        if actual != expected {
            t.Errorf("expected: %s \n received: %s \n", expected, actual)
        }
    })
}
